package models

type Category struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"not null"`
	Films []*Film `json:"films" gorm:"many2many:film_categories"`
}
