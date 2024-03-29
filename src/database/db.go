package database

import (
	"admin/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
    dsn = "admin:admin@tcp(db:3306)/ambassador?charset=utf8mb4&parseTime=True&loc=Local"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Product{})
}