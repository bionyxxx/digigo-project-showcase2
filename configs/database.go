package configs

import (
	"Challenge7/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOST     = "containers-us-west-21.railway.app"
	PORT     = "6898"
	USER     = "postgres"
	PASSWORD = "bOtYAzKvvCtEcoENk4kT"
	DBNAME   = "railway"
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
