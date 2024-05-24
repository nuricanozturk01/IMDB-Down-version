package entity

import (
	"github.com/google/uuid"
	"time"
)

type TVShow struct {
	ID           uuid.UUID   `json:"id" gorm:"type:char(36);primaryKey;"`
	Name         string      `json:"name" gorm:"type:varchar(80);"`
	Year         int         `json:"year" gorm:"type:int;"`
	Popularity   int         `json:"popularity" gorm:"type:int;"`
	AverageRate  float64     `json:"average_rate"`
	ClickCount   uint32      `json:"click_count" gorm:"type:int;"`
	EpisodeCount uint32      `json:"episode_count" gorm:"type:int;"`
	SeasonCount  uint32      `json:"season_count" gorm:"type:int;"`
	Photos       []Photo     `json:"photos" gorm:"foreignKey:OwnerID"`
	Trailers     []Trailer   `json:"trailers" gorm:"foreignKey:OwnerID"`
	Companies    []Company   `json:"companies" gorm:"foreignKey:OwnerID"`
	Celebs       []Celebrity `json:"celebs" gorm:"many2many:tvshow_celebs;"`
	Ratings      []Rating    `json:"ratings" gorm:"foreignKey:RateableID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
