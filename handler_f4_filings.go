package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeronimoLa/eagle/internal/edgar"
)

func (cfg *apiConfig) handlerF4Filings(w http.ResponseWriter, r *http.Request) {
	form := "4"
	tickerSymbol := r.URL.Query().Get("symbol")
	cik, err := cfg.db.GetCIKByTicker(context.Background(), tickerSymbol)
	if err != nil {
		log.Fatal("database call for GetCIKByTicker was not successful")
	}
	allRecentFilings, _ := cfg.edgar.StockFilings(cik)
	recentFilings, _ := edgar.StockRecentFilings(allRecentFilings)
	form4filings, _ := edgar.InsiderTransactions(cik, form, recentFilings)
	data, _ := cfg.edgar.ParseForm4Files(form4filings[:2])
	apiResp, _ := json.Marshal(data)
	w.Write(apiResp)

	// if tickerSymbol == "TSLA" {
	// data, _ := edgar.ParseForm4()
	// }
}
