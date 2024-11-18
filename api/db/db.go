package db

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"api/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	 dsn := "user:password@tcp(db:3306)/weather_db?charset=utf8mb4&parseTime=True&loc=Local"
	 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	 if err != nil {
		 log.Fatal(err)
	 }

	 db.AutoMigrate(&models.WeatherResponse{})
	 if err != nil {
		log.Fatal(err)
	 }

	 return db
}

func GetDB() *gorm.DB {
	return db
}