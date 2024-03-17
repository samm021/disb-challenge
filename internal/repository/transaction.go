package repository

import (
	"context"
	"disbursement-service/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransaction(connection *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: connection,
	}
}

func (t transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) string {
	t.db.Create(&transaction)

	return transaction.ReferenceId
}
