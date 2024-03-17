package domain

import (
	"context"
	"time"
)

type User struct {
	ID               uint          `json:"id" gorm:"primaryKey"`
	UserId           string        `json:"user_id" `
	Pin              int16         `json:"pin"`
	Balance          int64         `json:"balance"`
	AvailableBalance int64         `json:"available_balance"`
	Transactions     []Transaction `gorm:"foreignKey:UserId;references:UserId"`
	CreatedAt        time.Time
}

type UserRepository interface {
	FindByUserId(ctx context.Context, id string) (User, error)
}

type UserService interface {
	FindByUserId(ctx context.Context, userId string) (User, error)
}
