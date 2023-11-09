package repository

import (
	"context"
	"github.com/yaza-putu/online-test-dot/src/app/auth/entity"
	"github.com/yaza-putu/online-test-dot/src/database"
)

type userRepository struct {
	entity entity.User
}

func NewUserRepository() UserInterface {
	return &userRepository{
		entity: entity.User{},
	}
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	e := u.entity
	r := database.Instance.Where("email = ?", email).First(&e)
	return e, r.Error
}
