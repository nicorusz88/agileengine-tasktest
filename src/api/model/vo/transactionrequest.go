package vo

type TransactionRequest struct {
	 Type string `json:"type" binding:"required"`
	 Amount *float64 `json:"amount" binding:"required,numeric,gt=0"`
}
