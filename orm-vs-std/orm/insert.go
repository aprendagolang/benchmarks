package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func InsertORM(db *gorm.DB, data entities.Category) (entities.Category, error) {
	err := db.Create(&data).Error

	return data, err
}
