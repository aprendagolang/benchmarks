package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
func setup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bench_orm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.Category{})
	if err != nil {
		panic(err)
	}

	return db
}
