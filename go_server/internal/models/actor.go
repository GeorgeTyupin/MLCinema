package models

type Actor struct {
	ID          uint    `json:"id" gorm:"primaryKey"`                // Уникальный ID
	Name        string  `json:"name" gorm:"not null"`                // Имя актёра
	BirthYear   int     `json:"birth_year"`                          // Год рождения
	Nationality string  `json:"nationality"`                         // Национальность
	Films       []*Film `json:"filmid" gorm:"many2many:film_actors"` // внешний ключ на таблицу Film
}
