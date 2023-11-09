package repository

import (
	"context"
	"github.com/yaza-putu/online-test-dot/src/app/goods/entity"
	"github.com/yaza-putu/online-test-dot/src/utils"
)

type GoodsInterface interface {
	All(ctx context.Context, page int, take int) (utils.Pagination, error)
	Create(ctx context.Context, cat entity.Goods) (entity.Goods, error)
	Update(ctx context.Context, id string, user entity.Goods) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (entity.Goods, error)
}
