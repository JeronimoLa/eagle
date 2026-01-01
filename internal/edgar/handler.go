package edgar

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service *edgarService
}

func NewHandler(svc *edgarService) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) HandlerF4Filings(w http.ResponseWriter, r *http.Request) {
	h.service.IdentifyNewEntries()
	form := "4"
	tickerSymbol := r.URL.Query().Get("symbol")
	latestFilings, err := strconv.Atoi(r.URL.Query().Get("latest"))
	if err != nil {
		latestFilings = 20
		log.Printf("func HandlerF4Filings() defaulting to %d for latest filings\n", latestFilings)
	}
	cik, err := h.service.db.GetCIKByTicker(context.Background(), tickerSymbol)
	if err != nil {
		log.Println("database call for GetCIKByTicker was not successful", err)
		errorMessage := make(map[string]string)
		errorMessage["error"] = "Malformed or Unsupported ticker symbol"
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.Encode(errorMessage)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(buf.Bytes())
		return
	}
	allRecentFilings, _ := h.service.StockFilings(cik)
	recentFilings, _ := StockRecentFilings(allRecentFilings)
	form4filings, _ := InsiderTransactions(cik, form, recentFilings)
	data, _ := h.service.ParseForm4Files(form4filings[:latestFilings])
	apiResp, _ := json.Marshal(data)
	w.Write(apiResp)

}
