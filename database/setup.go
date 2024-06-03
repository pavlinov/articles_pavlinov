package database

import (
	"articles_pavlinov/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "test.db" // Значение по умолчанию, если переменная окружения не установлена
	}
	DB, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Preference{})
}
