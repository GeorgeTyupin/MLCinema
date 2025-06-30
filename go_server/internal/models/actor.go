package models

type Actor struct {
	ID          int    `json:"id"`          // Уникальный ID
	Name        string `json:"name"`        // Имя актёра
	BirthYear   int    `json:"birth_year"`  // Год рождения
	Nationality string `json:"nationality"` // Национальность
}
