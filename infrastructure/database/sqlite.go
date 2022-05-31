package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"omnisys.io/core/domain"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("omnisys.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(domain.Product{}, domain.Customer{}, domain.Order{}, domain.OrderItem{})

	return db
}
