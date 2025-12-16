package edgar

import (
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
	form := "4"
	tickerSymbol := r.URL.Query().Get("symbol")
	latestFilings, err := strconv.Atoi(r.URL.Query().Get("latest"))
	if err != nil {
		log.Println("func HandlerF4Filings() defaulting to 3 for latest filings")
		latestFilings = 3
	}
	cik, err := h.service.db.GetCIKByTicker(context.Background(), tickerSymbol)
	if err != nil {
		log.Fatal("database call for GetCIKByTicker was not successful")
	}
	allRecentFilings, _ := h.service.StockFilings(cik)
	recentFilings, _ := StockRecentFilings(allRecentFilings)
	form4filings, _ := InsiderTransactions(cik, form, recentFilings)
	data, _ := h.service.ParseForm4Files(form4filings[:latestFilings])
	apiResp, _ := json.Marshal(data)
	w.Write(apiResp)

}
