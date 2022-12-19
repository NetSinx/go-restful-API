package app

import (
	"fmt"
	"os"

	"github.com/NetSinx/go-restful-API/model/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var user domain.User
	var product domain.Product

	godotenv.Load()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASS"),
						os.Getenv("DB_HOST"),
						os.Getenv("DB_PORT"),
						os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)	
	}

	db.AutoMigrate(&user)
	db.AutoMigrate(&product)

	DB = db
}