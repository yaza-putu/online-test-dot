package migrations

import (
	"github.com/yaza-putu/online-test-dot/src/app/goods/entity"
	"github.com/yaza-putu/online-test-dot/src/database"
	"gorm.io/gorm"
)

/// please replace or change &entities.Name{}
/// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
// It will change existing column’s type if its size, precision, nullable changed.
// It WON’T delete unused columns to protect your data.

func init() {
	database.MigrationRegister(func(db *gorm.DB) error {
		return db.AutoMigrate(&entity.Goods{})
	}, func(db *gorm.DB) error {
		return db.Migrator().DropTable(&entity.Goods{})
	})
}
