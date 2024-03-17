package repository

import (
	"context"
	"disbursement-service/domain"

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

func (u userRepository) FindByUserId(ctx context.Context, userId string) (user domain.User, err error) {
	u.db.Where("user_id = ?", userId).First(&user)

	return
}
