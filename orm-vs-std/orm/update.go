package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func UpdateORM(db *gorm.DB, data entities.Category) error {
	return db.Save(&data).Error
}
