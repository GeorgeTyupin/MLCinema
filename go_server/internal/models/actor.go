package models

type Actor struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // Уникальный ID
	Name        string `json:"name" gorm:"not null"` // Имя актёра
	BirthYear   int    `json:"birth_year"`           // Год рождения
	Nationality string `json:"nationality"`          // Национальность
	FilmID      uint   `json:"filmid" `              // внешний ключ на таблицу Film
}
