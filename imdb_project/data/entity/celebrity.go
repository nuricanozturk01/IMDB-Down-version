package entity

import (
	"github.com/google/uuid"
)

type Celebrity struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key"`
	Name    string
	Movies  []Movie  `gorm:"many2many:movie_celebs;"`
	Photos  []Photo  `json:"photos" gorm:"foreignKey:OwnerID"`
	TVShows []TVShow `gorm:"many2many:tvshow_celebs;"`
}
