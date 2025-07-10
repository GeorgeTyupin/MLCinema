package models

type Film struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Title       string      `json:"title" gorm:"not null"`
	Year        uint        `json:"year"`
	Country     string      `json:"country"`
	Genre       string      `json:"genre"`
	ImagePath   string      `json:"imagePath"`
	Actors      []*Actor    `json:"actors" gorm:"many2many:film_actors"`
	Categories  []*Category `json:"categories" gorm:"many2many:film_categories"`
	Description string      `json:"description" gorm:"type:text"`
}
