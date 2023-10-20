package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
func setup() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        "host=localhost port=5432 user=postgres password=1234 dbname=benchmark sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.Category{})
	if err != nil {
		panic(err)
	}

	return db
}
