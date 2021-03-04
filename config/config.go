package config

// variabel DB yang merupakan pointer dari gorm.DB

import (
	"fmt"
	"log"
	"os"
	"todo/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading .env File")
	}
	dsn := "root:" + os.Getenv("PASS") + "@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.Project{}, &models.Label{}, &models.Task{})
	fmt.Println("DB COnnected")

	DB = db
}
