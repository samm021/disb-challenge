package repository

import (
	"context"
	"disbursement-service/domain"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(connection *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: connection,
	}
}

func (u userRepository) FindByUserId(ctx context.Context, userId string) (user domain.User) {
	u.db.Where("user_id = ?", userId).First(&user)

	return
}

func (u userRepository) UpdateAvailableBalance(ctx context.Context, user *domain.User, availableBalance decimal.Decimal) {
	u.db.Model(&user).Update("available_balance", availableBalance)
}

func (u userRepository) UpdateBalance(ctx context.Context, user *domain.User, balance decimal.Decimal) {
	u.db.Model(&user).Update("balance", balance)
}
