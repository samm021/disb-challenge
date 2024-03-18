package dto

type TransactionReq struct {
	Amount        float32 `json:"amount"`
	ChannelCode   string  `json:"channel_code"`
	Type          string  `json:"type"`
	Currency      string  `json:"currency"`
	Description   *string `json:"description"`
	AccountNumber string  `json:"account_number"`
	AccountType   string  `json:"account_type"`
}

type TransactionRes struct {
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
}
