package service

import (
	"context"
	"disbursement-service/domain"

	"github.com/shopspring/decimal"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) FindByUserId(ctx context.Context, userId string) domain.User {
	user := u.userRepository.FindByUserId(ctx, userId)
	return user
}

func (u userService) UpdateAvailableBalance(ctx context.Context, user *domain.User, availableBalance decimal.Decimal) {
	u.userRepository.UpdateAvailableBalance(ctx, user, availableBalance)
}

func (u userService) UpdateBalance(ctx context.Context, user *domain.User, balance decimal.Decimal) {
	u.userRepository.UpdateBalance(ctx, user, balance)
}
