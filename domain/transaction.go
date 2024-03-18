package domain

import (
	"context"
	"disbursement-service/dto"
	xenditDto "disbursement-service/dto/xendit"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID             uint            `json:"id" gorm:"primaryKey"`
	UserId         string          `json:"user_id"`
	ReferenceId    string          `json:"reference_id"`
	Amount         decimal.Decimal `json:"amount" gorm:"type:decimal(20,2)"`
	ChannelCode    string          `json:"channel_code"`
	Type           string          `json:"type"`
	Currency       string          `json:"currency"`
	Status         string          `json:"status"`
	Description    *string         `json:"description"`
	AccountNumber  string          `json:"account_number"`
	AccountType    string          `json:"account_type"`
	FailureCode    *string         `json:"failure_code"`
	IdempotencyKey string          `json:"idempotency_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *Transaction)
	FindByReferenceId(ctx context.Context, refId string) Transaction
	UpdateStatus(ctx context.Context, transaction *Transaction, status string, failureCode *string)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, req dto.TransactionReq) (dto.TransactionRes, error)
	UpdateTransactionStatus(ctx context.Context, req xenditDto.XenditPayoutCallbackReq, webhookId string) (dto.TransactionRes, error)
}
