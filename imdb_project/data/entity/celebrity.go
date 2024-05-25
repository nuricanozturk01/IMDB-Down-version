package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Celebrity struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key"`
	Name    string
	Movies  []Movie  `gorm:"many2many:movie_celebs;"`
	Photos  []Photo  `json:"photos" gorm:"polymorphic:Media;polymorphicValue:celebs"`
	TVShows []TVShow `gorm:"many2many:tvshow_celebs;"`
}

func (celebrity *Celebrity) BeforeCreate(tx *gorm.DB) (err error) {
	celebrity.ID = uuid.New()
	return
}
