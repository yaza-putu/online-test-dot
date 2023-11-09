package repository

import (
	"context"
	"github.com/yaza-putu/online-test-dot/src/app/category/entity"
	"github.com/yaza-putu/online-test-dot/src/database"
	"github.com/yaza-putu/online-test-dot/src/utils"
	"strings"
)

type categoryRepository struct {
	entity   entity.Category
	entities entity.Categories
}

func NewCategory() CategoryInterface {
	return &categoryRepository{
		entity:   entity.Category{},
		entities: entity.Categories{},
	}
}

func (c *categoryRepository) Create(ctx context.Context, cat entity.Category) (entity.Category, error) {
	cat.ID = utils.Uid(13)
	cat.Name = strings.ToTitle(cat.Name)
	r := database.Instance.WithContext(ctx).Create(&cat)

	return cat, r.Error
}

func (c *categoryRepository) Update(ctx context.Context, id string, cat entity.Category) error {
	cat.Name = strings.ToTitle(cat.Name)
	return database.Instance.WithContext(ctx).Where("id = ?", id).Updates(&cat).Error
}

func (c *categoryRepository) Delete(ctx context.Context, id string) error {
	return database.Instance.WithContext(ctx).Where("id = ?", id).Delete(&c.entity).Error
}

func (c *categoryRepository) FindById(ctx context.Context, id string) (entity.Category, error) {
	e := c.entity
	r := database.Instance.WithContext(ctx).Where("id = ?", id).First(&e)

	return e, r.Error
}

func (c *categoryRepository) All(ctx context.Context, page int, take int) (utils.Pagination, error) {
	e := c.entities
	var pagination utils.Pagination
	var totalRow int64

	r := database.Instance.WithContext(ctx).Model(&e)
	r.Count(&totalRow)
	r.Scopes(pagination.Paginate(page, take))
	r.Find(&e)

	pagination.Rows = e
	pagination.CalculatePage(float64(totalRow))

	return pagination, r.Error
}
