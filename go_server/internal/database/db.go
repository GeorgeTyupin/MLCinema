package database

import (
	"fmt"
	"log"
	"os"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	DB = db

	// DB.Migrator().DropTable(&models.Film{}, &models.Actor{}, &models.Category{}, "film_categories", "film_actors") // Удаление старых таблиц для миграций

	if err := DB.AutoMigrate(&models.Actor{}, &models.Film{}, &models.Category{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	// SeedTestData()
}
