package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"omnisys.io/core/domain"
)

func ConnectDB() *gorm.DB {
	// sqlite
	db, err := gorm.Open(sqlite.Open("omnisys.db"), &gorm.Config{})

	// // mysql
	// dsn := "root:@tcp(127.0.0.1:3306)/omnimysql?parseTime=true"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(domain.Product{}, domain.Customer{}, domain.Order{}, domain.OrderItem{})

	return db
}
