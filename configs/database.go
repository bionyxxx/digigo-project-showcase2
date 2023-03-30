package configs

import (
	"Challenge7/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "challenge7"
	DB       *gorm.DB
	err      error
)

func GetDB() *gorm.DB {
	return DB
}

func StartDBConnection() {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DBNAME, PASSWORD)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err := DB.Debug().AutoMigrate(&models.Book{})
	if err != nil {
		panic(err)
	}
}
