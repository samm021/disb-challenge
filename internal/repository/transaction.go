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

func (t transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) {
	t.db.Create(&transaction)
}

func (t transactionRepository) FindByReferenceId(ctx context.Context, referenceId string) (transaction domain.Transaction) {
	t.db.Where("reference_id = ?", referenceId).First(&transaction)
	return
}

func (t transactionRepository) UpdateStatus(ctx context.Context, transaction *domain.Transaction, status string, failureCode *string) {
	if failureCode == nil {
		t.db.Model(&transaction).Updates(domain.Transaction{Status: status, FailureCode: failureCode})
		return
	}

	t.db.Model(&transaction).Update("status", status)
}
