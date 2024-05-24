package entity

import (
	"github.com/google/uuid"
	"time"
)

type Movie struct {
	ID          uuid.UUID   `json:"id" gorm:"type:char(36);primaryKey;"`
	Name        string      `json:"name" gorm:"type:varchar(80);"`
	AverageRate float64     `json:"average_rate" gorm:"type:float; default:0;"`
	Year        uint        `json:"year" gorm:"type:int;"`
	Popularity  uint        `json:"popularity" gorm:"type:int;default:0;"`
	ClickCount  uint        `json:"click_count" gorm:"type:int; default:0;"`
	Trailers    []Trailer   `json:"trailers" gorm:"foreignKey:OwnerID"`
	Companies   []Company   `json:"companies" gorm:"foreignKey:OwnerID"`
	Celebs      []Celebrity `json:"celebs" gorm:"many2many:movie_celebs;"`
	Ratings     []Rating    `json:"ratings" gorm:"foreignKey:RateableID"`
	Photos      []Photo     `json:"photos" gorm:"foreignKey:OwnerID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
