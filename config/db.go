package config

import (
	"fmt"
	"log"
	"studynotes/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/study_api?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to MySQL database")

	err = DB.AutoMigrate(
		&models.Users{},
		&models.Topic{},
	)
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}
	fmt.Println("✅ Database migration completed")

	return DB // ✅ tambahkan ini
}
