package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func ListORM(db *gorm.DB) ([]entities.Category, error) {
	var categories []entities.Category
	err := db.Find(&categories).Error
	return categories, err
}
