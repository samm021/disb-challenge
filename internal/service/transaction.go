package service

import (
	"context"
	"disbursement-service/domain"
	"disbursement-service/dto"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
}

func NewTransaction(transactionRepository domain.TransactionRepository) domain.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
	}
}

func (t transactionService) CreateTransaction(ctx context.Context, req dto.TransactionReq) (dto.TransactionRes, error) {
	return dto.TransactionRes{}, nil
}
