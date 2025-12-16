package edgar

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jeronimoLa/eagle/internal/client"
	"github.com/jeronimoLa/eagle/internal/database"
)

type edgarService struct {
	client *client.Config
	db     *database.Queries
}

type TransactionMapping struct {
	AccessionNumber string
	Url             string
}

func NewService(client *client.Config, db *database.Queries) *edgarService {
	return &edgarService{
		client: client,
		db:     db,
	}
}

func (cfg *edgarService) StockFilings(CIK string) (*FilingSubmissions, error) {
	submissionURL := fmt.Sprintf("https://data.sec.gov/submissions/%s.json", CIK)
	req, err := http.NewRequest(http.MethodGet, submissionURL, nil)
	if err != nil {
		log.Fatalf("Unable to create the http request")
	}

	req.Header.Set("User-Agent", cfg.client.UserAgent)

	resp, err := cfg.client.HTTPClient.Do(req)
	if err != nil {
		log.Fatalf("Unable to make an HTTP request to %s", submissionURL)
	}
	defer resp.Body.Close()
	var fs FilingSubmissions
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &fs)
	return &fs, nil
}

func StockRecentFilings(fs *FilingSubmissions) ([]map[string]interface{}, error) {
	recentJSON, _ := json.Marshal(fs.Filings.Recent)
	var recentMap map[string]interface{}
	json.Unmarshal(recentJSON, &recentMap)

	entries := make([]map[string]interface{}, 1000)
	for i := range entries {
		entries[i] = make(map[string]interface{})
	}
	for k, v := range recentMap {
		list, ok := v.([]interface{})
		if !ok {
			fmt.Println("Skipping non-list key", k)
		}

		for j, entry := range entries {
			entry[k] = list[j]
		}
	}
	return entries, nil
}

func InsiderTransactions(cik string, form string, filings []map[string]interface{}) ([]TransactionMapping, error) {
	// https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-69
	var form4Filings []TransactionMapping
	cikNum := strings.TrimLeft(cik, "CIK0")
	for _, filing := range filings {
		if filing["form"].(string) == form {
			accessionNumber := strings.Replace(filing["accessionNumber"].(string), "-", "", -1)
			primaryDocument := strings.Split(filing["primaryDocument"].(string), "/")
			form4Url := fmt.Sprintf("https://www.sec.gov/Archives/edgar/data/%s/%s/%s", cikNum, accessionNumber, primaryDocument[1])
			form4Filings = append(form4Filings, TransactionMapping{AccessionNumber: accessionNumber, Url: form4Url})
		}
	}
	fmt.Printf("number of form 4 filings out of 1000: %d\n", len(form4Filings))
	return form4Filings, nil
}

func (cfg *edgarService) GetForm4Filings(form4URL string) (*XMLForm4, error) {
	items := strings.Split(form4URL, "/")
	accessionNumber := items[len(items)-2]
	req, err := http.NewRequest(http.MethodGet, form4URL, nil)
	if err != nil {
		log.Fatalf("Unable to create the http request")
	}
	req.Header.Set("User-Agent", cfg.client.UserAgent)
	resp, err := cfg.client.HTTPClient.Do(req)
	if err != nil {
		log.Printf("call to %s was unsuccessful", form4URL)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Unable to make an HTTP request to %s", form4URL)
	}
	body, _ := io.ReadAll(resp.Body)
	var form4filing XMLForm4
	xml.Unmarshal(body, &form4filing)

	cfg.db.InsertF4FilingRecord(context.Background(), database.InsertF4FilingRecordParams{
		AccessionNumber:     accessionNumber,
		DocumentType:        form4filing.DocumentType,
		PeriodOfReport:      form4filing.PeriodOfReport,
		IssuerCik:           form4filing.IssuerCIK,
		Issuertradingsymbol: form4filing.IssuerTradingSymbol,
		RptOwnerCik:         form4filing.ReportingOwner.ReportingOwnerId.RptOwnerCik,
		RptOwnerName:        form4filing.ReportingOwner.ReportingOwnerId.RptOwnerName,
		OfficerTitle:        form4filing.ReportingOwner.ReportingOwnerRelationship.OfficerTitle,
		CreatedAt:           time.Now(),
		Form4Url:            form4URL,
	})
	return &form4filing, nil
}
