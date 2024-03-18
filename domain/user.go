package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID               uint            `json:"id" gorm:"primaryKey"`
	UserId           string          `json:"user_id" `
	Pin              int16           `json:"pin"`
	Balance          decimal.Decimal `json:"balance" gorm:"type:decimal(20,2)"`
	AvailableBalance decimal.Decimal `json:"available_balance" gorm:"type:decimal(20,2)"`
	Transactions     []Transaction   `gorm:"foreignKey:UserId;references:UserId"`
	CreatedAt        time.Time
}

type UserRepository interface {
	FindByUserId(ctx context.Context, id string) User
	UpdateAvailableBalance(ctx context.Context, user *User, availableBalance decimal.Decimal)
	UpdateBalance(ctx context.Context, user *User, balance decimal.Decimal)
}

type UserService interface {
	FindByUserId(ctx context.Context, userId string) User
	UpdateAvailableBalance(ctx context.Context, user *User, availableBalance decimal.Decimal)
	UpdateBalance(ctx context.Context, user *User, balance decimal.Decimal)
}
