package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func DeleteORM(db *gorm.DB, id int64) error {
	return db.Delete(&entities.Category{ID: id}).Error
}
