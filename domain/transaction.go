package domain

import (
	"context"
	"disbursement-service/dto"
	"time"
)

type Transaction struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	UserId        string  `json:"user_id"`
	ReferenceId   string  `json:"reference_id"`
	Amount        int64   `json:"amount"`
	ChannelCode   string  `json:"channel_code"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
	Description   *string `json:"description"`
	AccountNumber string  `json:"account_number"`
	AccountType   string  `json:"account_type"`
	FailureCode   *string `json:"failure_code"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *Transaction) string
	// FindByReferenceId(ctx context.Context, refId string) (Transaction, error)
	// UpdateStatus(ctx context.Context, refId string, status string) (Transaction, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, req dto.TransactionReq) (dto.TransactionRes, error)
}
