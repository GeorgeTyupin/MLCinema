package database

import (
	"log"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
)

func SeedTestData() {
	var count int64
	DB.Model(&models.Film{}).Count(&count)
	if count > 0 {
		return // Уже есть фильмы — сидер не нужен
	}

	film := models.Film{
		Title:       "Интерстеллар",
		Year:        2014,
		Country:     "США",
		Genre:       "Фантастика",
		Description: "Будущее Земли под угрозой. Команда астронавтов исследует другие галактики.",
		Actors: []*models.Actor{
			{Name: "Мэттью МакКонахи"},
			{Name: "Энн Хэтэуэй"},
			{Name: "Джессика Честейн"},
		},
	}

	if err := DB.Create(&film).Error; err != nil {
		log.Printf("Ошибка при создании тестовых данных: %v", err)
	}
}
