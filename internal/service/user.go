package service

import (
	"context"
	"disbursement-service/domain"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) FindByUserId(ctx context.Context, userId string) (domain.User, error) {
	user, err := u.userRepository.FindByUserId(ctx, userId)
	return user, err
}
