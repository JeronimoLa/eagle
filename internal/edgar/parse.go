package edgar

import (
	"github.com/Rhymond/go-money"
)

func (cfg *edgarService) ParseForm4Files(form4files []TransactionMapping) (map[string]ValidTransactionSignal, error) {
	// Non-derivative = actual shares of stock
	// Derivative = options, RSUs, convertible securities (confusing, not sentiment)
	data := make(map[string]ValidTransactionSignal)
	for _, file := range form4files {
		res, _ := cfg.GetForm4Filings(file.Url)
		NumberOfDerivativeTransactions := len(res.DerivativeTable.DerivativeTransaction)
		if NumberOfDerivativeTransactions >= 1 {
			continue
		}

		sharesBought, sharesSold := 0.0, 0.0
		MoneyIn, MoneyOut := money.New(0, money.USD), money.New(0, money.USD)
		var obj []Transactions
		for _, transac := range res.NonDerivativeTable.NonDerivativeTransaction {
			transactionPrice := money.NewFromFloat(transac.TransactionPricePerShare, money.USD)
			totalAmount := transactionPrice.Multiply(int64(transac.TransactionShares))

			if transac.TransactionAcquiredDisposedCode == "A" {
				MoneyIn, _ = MoneyIn.Add(totalAmount)
				sharesBought += transac.TransactionShares
			}

			if transac.TransactionAcquiredDisposedCode == "D" {
				MoneyOut, _ = MoneyOut.Add(totalAmount)
				sharesSold += transac.TransactionShares
			}
			pricePerShare := money.NewFromFloat(transac.TransactionPricePerShare, money.USD)
			transactionTotal := pricePerShare.Multiply(int64(transac.TransactionShares))

			obj = append(obj, Transactions{
				SecurityTitle:                   transac.SecurityTitle,
				RptOwner:                        res.ReportingOwner.ReportingOwnerId.RptOwnerName,
				SharesBought:                    sharesBought,
				SharesSold:                      sharesSold,
				TransactionDate:                 transac.TransactionDate,
				TransactionCode:                 transac.TransactionCoding,
				TransactionAcquiredDisposedCode: transac.TransactionAcquiredDisposedCode,
				TransactionShares:               transac.TransactionShares,
				TransactionPricePerShare:        pricePerShare.Display(),
				SharesOwnedFollowingTransaction: transac.PostTransactionAmounts,
				TransactionTotal:                transactionTotal.Display(),
			})
		}
		data[file.AccessionNumber] = ValidTransactionSignal{
			CIK:          res.ReportingOwner.ReportingOwnerId.RptOwnerCik,
			DocumentType: res.DocumentType,
			Data:         obj,
		}
	}
	return data, nil

}
