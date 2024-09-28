package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	
	dsn := "rastogi:Rastogi@2004@tcp(127.0.0.1:3306)/bookLine?charset=utf8mb4&parseTime=True&loc=Local"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
