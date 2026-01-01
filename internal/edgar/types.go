package edgar

import (
	"time"
)

// JSON model for ticker specifc submissions
type FilingSubmissions struct {
	Cik                             string   `json:"cik"`
	EntityType                      string   `json:"entityType"`
	Sic                             string   `json:"sic"`
	SicDescription                  string   `json:"sicDescription"`
	OwnerOrg                        string   `json:"ownerOrg"`
	Name                            string   `json:"name"`
	Tickers                         []string `json:"tickers"`
	Exchanges                       []string `json:"exchanges"`
	Ein                             string   `json:"ein"`
	Category                        string   `json:"category"`
	StateOfIncorporation            string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription string   `json:"stateOfIncorporationDescription"`
	FormerNames                     []struct {
		Name string    `json:"name"`
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"formerNames"`
	Filings struct {
		Recent struct {
			AccessionNumber       []string    `json:"accessionNumber"`
			FilingDate            []string    `json:"filingDate"`
			ReportDate            []string    `json:"reportDate"`
			AcceptanceDateTime    []time.Time `json:"acceptanceDateTime"`
			Act                   []string    `json:"act"`
			Form                  []string    `json:"form"`
			FileNumber            []string    `json:"fileNumber"`
			FilmNumber            []string    `json:"filmNumber"`
			Items                 []string    `json:"items"`
			CoreType              []string    `json:"core_type"`
			Size                  []int       `json:"size"`
			IsXBRL                []int       `json:"isXBRL"`
			IsInlineXBRL          []int       `json:"isInlineXBRL"`
			PrimaryDocument       []string    `json:"primaryDocument"`
			PrimaryDocDescription []string    `json:"primaryDocDescription"`
		} `json:"recent"`
		Files []struct {
			Name        string `json:"name"`
			FilingCount int    `json:"filingCount"`
			FilingFrom  string `json:"filingFrom"`
			FilingTo    string `json:"filingTo"`
		} `json:"files"`
	} `json:"filings"`
}

// FORM4 XML model
type XMLForm4 struct {
	DocumentType        string             `xml:"documentType"`
	PeriodOfReport      string             `xml:"periodOfReport"`
	IssuerCIK           string             `xml:"issuer>issuerCik"`
	IssuerTradingSymbol string             `xml:"issuer>issuerTradingSymbol"`
	ReportingOwner      ReportingOwner     `xml:"reportingOwner"`
	NonDerivativeTable  NonDerivativeTable `xml:"nonDerivativeTable"`
	DerivativeTable     DerivativeTable    `xml:"derivativeTable"`
}

type ReportingOwner struct {
	ReportingOwnerId           ReportingOwnerId           `xml:"reportingOwnerId"`
	ReportingOwnerRelationship ReportingOwnerRelationship `xml:"ReportingOwnerRelationship"`
}

type ReportingOwnerId struct {
	RptOwnerCik  string `xml:"rptOwnerCik"`
	RptOwnerName string `xml:"rptOwnerName"`
}

type ReportingOwnerRelationship struct {
	IsDirector        int    `xml:"isDirector"`
	IsOfficer         int    `xml:"isOfficer"`
	IsTenPercentOwner int    `xml:"isTenPercentOwner"`
	IsOther           int    `xml:"isOther"`
	OfficerTitle      string `xml:"officerTitle"`
}

type NonDerivativeTable struct {
	NonDerivativeTransaction []NonDerivativeTransaction `xml:"nonDerivativeTransaction"`
}

type NonDerivativeTransaction struct {
	SecurityTitle                   string  `xml:"securityTitle>value"`
	TransactionDate                 string  `xml:"transactionDate>value"`
	TransactionCoding               string  `xml:"transactionCoding>transactionCode"`
	TransactionAcquiredDisposedCode string  `xml:"transactionAmounts>transactionAcquiredDisposedCode>value"`
	TransactionShares               float64 `xml:"transactionAmounts>transactionShares>value"`
	TransactionPricePerShare        float64 `xml:"transactionAmounts>transactionPricePerShare>value"`
	PostTransactionAmounts          string  `xml:"postTransactionAmounts>sharesOwnedFollowingTransaction>value"`
}

type DerivativeTransaction struct {
	SecurityTitle string `xml:"securityTitle"`
}

type DerivativeTable struct {
	DerivativeTransaction []DerivativeTransaction `xml:"derivativeTransaction"`
}

type Transactions struct {
	SecurityTitle                   string  `json:"security_title"`
	RptOwner                        string  `json:"rpt_owner"`
	SharesBought                    float64 `json:"shares_bought"`
	SharesSold                      float64 `json:"shares_sold"`
	TransactionDate                 string  `json:"transaction_date"`
	TransactionCode                 string  `json:"transaction_code"`
	TransactionAcquiredDisposedCode string  `json:"acquired_disposed"`
	TransactionShares               float64 `json:"transaction_shares"`
	TransactionPricePerShare        string  `json:"share_price"`
	SharesOwnedFollowingTransaction string  `json:"remaining_shares"`
	TransactionTotal                string  `json:"transaction_total"`
}

type ValidTransactionSignal struct {
	CIK          string         `json:"cik"`
	DocumentType string         `json:"document_type"`
	Data         []Transactions `json:"data"`
}
