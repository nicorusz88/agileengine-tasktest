package vo

import "time"

type TransactionResponse struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	EffectiveDate time.Time `json:"effective_date"`
}