package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/go_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot conect to database")
	}

	fmt.Println("connect to database")
}

func CloseDatabaseConnection(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	database.Close()
}
