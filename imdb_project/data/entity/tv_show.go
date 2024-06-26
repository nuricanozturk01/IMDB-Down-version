package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TVShow struct {
	ID           uuid.UUID   `json:"id" gorm:"type:char(36);primaryKey;"`
	Name         string      `json:"name" gorm:"type:varchar(80);"`
	Year         int         `json:"year" gorm:"type:int;"`
	Description  string      `json:"description" gorm:"type:text;"`
	Popularity   int         `json:"popularity" gorm:"type:int;"`
	AverageRate  float64     `json:"average_rate"`
	ClickCount   uint32      `json:"click_count" gorm:"type:int;"`
	EpisodeCount uint32      `json:"episode_count" gorm:"type:int;"`
	SeasonCount  uint32      `json:"season_count" gorm:"type:int;"`
	Photos       []Photo     `json:"photos" gorm:"polymorphic:Media;polymorphicValue:tv_shows"`
	Trailers     []Trailer   `json:"trailers" gorm:"polymorphic:Media;polymorphicValue:tv_shows"`
	Companies    []Company   `json:"companies" gorm:"polymorphic:Media;polymorphicValue:tv_shows"`
	Celebs       []Celebrity `json:"celebs" gorm:"many2many:tvshow_celebs;"`
	Likes        []Like      `gorm:"polymorphic:Media;polymorphicValue:tv_shows"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (tvs *TVShow) BeforeCreate(tx *gorm.DB) (err error) {
	if tvs.ID == uuid.Nil {
		tvs.ID = uuid.New()
	}
	return
}
func (tvs *TVShow) GetID() uuid.UUID {
	return tvs.ID
}

func (tvs *TVShow) GetName() string {
	return tvs.Name
}
