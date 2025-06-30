package models

type Film struct {
	ID          int     `json:"id"`          // Уникальный ID
	Title       string  `json:"title"`       // Название фильма
	Year        int     `json:"year"`        // Год выпуска
	Country     string  `json:"country"`     // Страна
	Genre       string  `json:"genre"`       // Жанр
	Actors      []Actor `json:"actors"`      // Список актёров
	Description string  `json:"description"` // Краткое описание
}
