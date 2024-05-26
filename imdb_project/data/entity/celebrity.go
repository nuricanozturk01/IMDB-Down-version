package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Celebrity struct {
	ID      uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name    string    `gorm:"type:varchar(80);" json:"name"`
	Movies  []Movie   `gorm:"many2many:movie_celebs;" json:"movies"`
	Photos  []Photo   `json:"photos" gorm:"polymorphic:Media;polymorphicValue:celebs"`
	TVShows []TVShow  `gorm:"many2many:tvshow_celebs;" json:"tv_shows"`
}

func (celebrity *Celebrity) BeforeCreate(tx *gorm.DB) (err error) {
	if celebrity.ID == uuid.Nil {
		celebrity.ID = uuid.New()
	}
	return
}
