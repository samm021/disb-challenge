package service

import (
	"context"
	"disbursement-service/domain"
	"disbursement-service/internal/config"

	"github.com/shopspring/decimal"
	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/payout"
)

type xenditService struct {
	xenditClient *xendit.APIClient
}

func NewXendit(config *config.Config) domain.XenditService {
	xenditClient := xendit.NewClient(config.Xendit.XApiKey)

	return &xenditService{xenditClient: xenditClient}
}

func (x xenditService) CreateDisbursementPayout(ctx context.Context, transactionReq *domain.Transaction) (domain.XenditPayoutStatus, error) {
	amount, _ := transactionReq.Amount.BigFloat().Float32()

	createPayoutRequest := *payout.NewCreatePayoutRequest(transactionReq.ReferenceId, transactionReq.ChannelCode, *payout.NewDigitalPayoutChannelProperties(transactionReq.AccountNumber), amount, transactionReq.Currency)

	resp, _, err := x.xenditClient.PayoutApi.CreatePayout(ctx).
		IdempotencyKey(transactionReq.IdempotencyKey).
		CreatePayoutRequest(createPayoutRequest).
		Execute()

	if err != nil {
		return domain.XenditPayoutStatus{}, err
	}

	return domain.XenditPayoutStatus{
		Status: resp.Payout.GetStatus(),
	}, nil
}

func (x xenditService) GetChannels(ctx context.Context) ([]domain.XenditChannel, error) {
	res, _, err := x.xenditClient.PayoutApi.GetPayoutChannels(ctx).Execute()
	if err != nil {
		return nil, err
	}

	channels := make([]domain.XenditChannel, 0, len(res))

	for i := 0; i < len(res); i++ {
		channel := domain.XenditChannel{
			MinAmount:       decimal.NewFromFloat32(res[i].AmountLimits.GetMinimum()),
			MaxAmount:       decimal.NewFromFloat32(res[i].AmountLimits.GetMaximum()),
			MinIncrement:    decimal.NewFromFloat32(res[i].AmountLimits.GetMinimumIncrement()),
			ChannelCode:     res[i].GetChannelCode(),
			ChannelCategory: res[i].GetChannelCategory().String(),
			Currency:        res[i].GetCurrency(),
			ChannelName:     res[i].GetChannelName(),
		}
		channels = append(channels, channel)
	}

	return channels, nil
}
