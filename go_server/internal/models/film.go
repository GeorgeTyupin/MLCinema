package models

type Film struct {
	ID          uint    `json:"id" gorm:"primaryKey"`            // Уникальный ID
	Title       string  `json:"title" gorm:"not null"`           // Название фильма
	Year        uint    `json:"year"`                            // Год выпуска
	Country     string  `json:"country"`                         // Страна
	Genre       string  `json:"genre"`                           // Жанр
	Actors      []Actor `json:"actors" gorm:"foreignKey:FilmID"` // Список актёров
	Description string  `json:"description"`                     // Краткое описание
}
