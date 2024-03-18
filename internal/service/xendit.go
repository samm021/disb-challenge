package service

import (
	"context"
	"disbursement-service/domain"
	"disbursement-service/internal/config"
	"log"

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

	// BLOCKER: Error account_type & account_holder_name is required but upon creating proper struct still rejected
	// last error:
	// 2024/03/18 14:17:15 ERROR: map[errorCode:API_VALIDATION_ERROR errorMessage:There was an error with the format submitted to the server. rawResponse:map[error_code:API_VALIDATION_ERROR errors:[map[message:"channel_properties.account_type" is not allowed path:channel_properties, account_type]] message:There was an error with the format submitted to the server.] status:400]

	accountHolderName := "TEST"
	payoutProperty := payout.DigitalPayoutChannelProperties{
		AccountNumber:     transactionReq.AccountNumber,
		AccountHolderName: *payout.NewNullableString(&accountHolderName),
		AccountType:       payout.CHANNELACCOUNTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK.Ptr(),
	}

	// createPayoutRequest := *payout.NewCreatePayoutRequest(transactionReq.ReferenceId, transactionReq.ChannelCode, *payout.NewDigitalPayoutChannelProperties(transactionReq.AccountNumber), amount, transactionReq.Currency)
	createPayoutRequest := *payout.NewCreatePayoutRequest(transactionReq.ReferenceId, transactionReq.ChannelCode, payoutProperty, amount, transactionReq.Currency)

	// TODO: use rest http call without xendit client (?)
	resp, _, err := x.xenditClient.PayoutApi.CreatePayout(ctx).
		IdempotencyKey(transactionReq.IdempotencyKey).
		CreatePayoutRequest(createPayoutRequest).
		Execute()

	if err != nil {
		log.Printf("ERROR: %s", err.FullError())
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
