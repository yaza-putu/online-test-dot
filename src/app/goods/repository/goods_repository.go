package repository

import (
	"context"
	catEntity "github.com/yaza-putu/online-test-dot/src/app/category/entity"
	"github.com/yaza-putu/online-test-dot/src/app/goods/entity"
	"github.com/yaza-putu/online-test-dot/src/database"
	"github.com/yaza-putu/online-test-dot/src/logger"
	"github.com/yaza-putu/online-test-dot/src/utils"
	"strings"
)

type goodsRepository struct {
	entity   entity.Goods
	entities entity.AllGoods
}

func NewGoods() GoodsInterface {
	return &goodsRepository{
		entity:   entity.Goods{},
		entities: entity.AllGoods{},
	}
}

// Create repository use database transaction
func (c *goodsRepository) Create(ctx context.Context, gds entity.Goods) (entity.Goods, error) {
	db := database.Instance
	// start transaction
	db.Begin()

	gds.ID = utils.Uid(13)
	gds.Name = strings.ToTitle(gds.Name)

	cat := catEntity.Category{}

	cf := db.Where("id = ?", gds.CategoryId).First(&cat)
	logger.New(cf.Error, logger.SetType(logger.ERROR))

	// check if category not exist and let's create
	if cat.ID == "" {
		catId := utils.Uid(13)
		catR := catEntity.Category{Name: strings.ToTitle(gds.CategoryId), ID: catId}
		cr := db.Create(&catR)
		if cr.Error != nil {
			db.Rollback()
		}
		gds.CategoryId = catId
		gds.Category = catR
	}

	r := db.WithContext(ctx).Create(&gds)

	// rollback if error
	if r.Error != nil {
		db.Rollback()
	}

	// if no error
	db.Commit()

	return gds, r.Error
}

func (c *goodsRepository) Update(ctx context.Context, id string, gds entity.Goods) error {
	gds.Name = strings.ToTitle(gds.Name)
	return database.Instance.WithContext(ctx).Where("id = ?", id).Updates(&gds).Error
}

func (c *goodsRepository) Delete(ctx context.Context, id string) error {
	return database.Instance.WithContext(ctx).Where("id = ?", id).Delete(&c.entity).Error
}

func (c *goodsRepository) FindById(ctx context.Context, id string) (entity.Goods, error) {
	e := c.entity
	r := database.Instance.WithContext(ctx).Preload("Category").Where("id = ?", id).First(&e)

	return e, r.Error
}

func (c *goodsRepository) All(ctx context.Context, page int, take int) (utils.Pagination, error) {
	e := c.entities
	var pagination utils.Pagination
	var totalRow int64

	r := database.Instance.WithContext(ctx).Model(&e)
	r.Count(&totalRow)
	r.Scopes(pagination.Paginate(page, take))
	r.Preload("Category").Find(&e)

	pagination.Rows = e
	pagination.CalculatePage(float64(totalRow))

	return pagination, r.Error
}
