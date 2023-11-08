package repository

import (
	"context"
	"github.com/yaza-putu/online-test-dot/src/app/auth/entity"
)

type UserInterface interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, id string, user entity.User) error
	Delete(ctx context.Context, id string) error
	FindOrFail(ctx context.Context, id string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}
