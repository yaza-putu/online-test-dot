package entity

import (
	"github.com/yaza-putu/online-test-dot/src/app/category/entity"
	"time"
)

type Goods struct {
	ID         string `gorm:"primaryKey;type:char(36)" json:"id"`
	CategoryId string `gorm:"type:char(36)" json:"category_id"`
	Category   entity.Category
	Name       string `gorm:"name" json:"name"`
	Stock      uint   `gorm:"stock;default:0" json:"stock"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AllGoods []Goods
