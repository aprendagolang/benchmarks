package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func GetORM(db *gorm.DB, id int64) (*entities.Category, error) {
	var category entities.Category
	err := db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
