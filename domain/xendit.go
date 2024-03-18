package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type XenditPayoutStatus struct {
	Status string `json:"status"`
}

type XenditChannel struct {
	ID              uint            `json:"id"`
	MinAmount       decimal.Decimal `json:"min_amount" gorm:"type:decimal(20,2)"`
	MaxAmount       decimal.Decimal `json:"max_amount" gorm:"type:decimal(20,2)"`
	MinIncrement    decimal.Decimal `json:"min_increment" gorm:"type:decimal(20,2)"`
	ChannelCode     string          `json:"channel_code"`
	ChannelCategory string          `json:"channel_category"`
	ChannelName     string          `json:"channel_name"`
	Currency        string          `json:"currency"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type XenditService interface {
	CreateDisbursementPayout(ctx context.Context, transaction *Transaction) (XenditPayoutStatus, error)
	GetChannels(ctx context.Context) ([]XenditChannel, error)
}
