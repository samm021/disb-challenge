package service

import (
	"context"
	"crypto/rand"
	"disbursement-service/domain"
	"disbursement-service/dto"
	xenditDto "disbursement-service/dto/xendit"
	"disbursement-service/internal/util"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
	userService           domain.UserService
	xenditService         domain.XenditService
}

func NewTransaction(transactionRepository domain.TransactionRepository, userService domain.UserService, xenditService domain.XenditService) domain.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		userService:           userService,
		xenditService:         xenditService,
	}
}

func (t transactionService) CreateTransaction(ctx context.Context, req dto.TransactionReq) (dto.TransactionRes, error) {
	transactionType, status, userId := req.Type, "INITIATED", ctx.Value("x-user").(dto.UserData).UserId
	// validate user value from context
	if userId == "" {
		return dto.TransactionRes{}, errors.New("userId")
	}
	// get actual user
	user := t.userService.FindByUserId(ctx, userId)

	// validate amount & available balance
	if req.Amount == 0 {
		return dto.TransactionRes{}, errors.New("Amount")
	}
	balance, _ := user.AvailableBalance.BigFloat().Float32()
	if req.Amount > balance {
		return dto.TransactionRes{}, errors.New("Balance")
	}

	// validate with available xendit channels
	channels, err := t.xenditService.GetChannels(ctx)
	if err != nil {
		panic("")
	}

	channel := util.Find(channels, func(c *domain.XenditChannel) bool {
		return c.ChannelCode == req.ChannelCode && c.Currency == req.Currency
	})
	if channel == nil {
		return dto.TransactionRes{}, errors.New("Channel")
	}
	if channel.MaxAmount.LessThan(decimal.NewFromFloat32(req.Amount)) {
		return dto.TransactionRes{}, errors.New("max amount")
	}
	if channel.MinAmount.GreaterThan(decimal.NewFromFloat32(req.Amount)) {
		return dto.TransactionRes{}, errors.New("min amount")
	}

	transactionAmount := decimal.NewFromFloat32(req.Amount)

	// create random string for reference id & idempotency key
	referenceId := make([]byte, 10)
	rand.Read(referenceId)
	idempotencyKey := make([]byte, 20)
	rand.Read(idempotencyKey)

	transaction := domain.Transaction{
		UserId:         user.UserId,
		Amount:         transactionAmount,
		ChannelCode:    req.ChannelCode,
		Currency:       req.Currency,
		Type:           req.Type,
		Description:    req.Description,
		AccountNumber:  req.AccountNumber,
		AccountType:    req.AccountType,
		ReferenceId:    fmt.Sprintf("trx-%X", referenceId),
		Status:         status,
		IdempotencyKey: fmt.Sprintf("key-%X", idempotencyKey),
	}

	// TODO: proper DB transaction starting here & ends after successful call to PG
	t.transactionRepository.Create(ctx, &transaction)
	t.userService.UpdateAvailableBalance(ctx, &user, user.AvailableBalance.Sub(transactionAmount))

	if transactionType != "DISBURSEMENT" {
		return dto.TransactionRes{}, errors.New("payment type not supported")
	}

	// TODO: fix create payout

	res, err := t.xenditService.CreateDisbursementPayout(ctx, &transaction)
	if err != nil {
		return dto.TransactionRes{}, errors.New("error when creating disbursement")
	}
	status = res.Status

	// get updated transaction just in case callback already update the transaction
	updatedTransaction := t.transactionRepository.FindByReferenceId(ctx, transaction.ReferenceId)
	if updatedTransaction.Status != "INITIATED" {
		status = updatedTransaction.Status
	}

	return dto.TransactionRes{
		ReferenceId: transaction.ReferenceId,
		Status:      status,
	}, nil
}

func (t transactionService) UpdateTransactionStatus(ctx context.Context, req xenditDto.XenditPayoutCallbackReq, webhookId string) (dto.TransactionRes, error) {
	return dto.TransactionRes{}, nil
}
