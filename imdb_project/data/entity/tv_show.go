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

func (tvShow *TVShow) BeforeCreate(tx *gorm.DB) (err error) {
	tvShow.ID = uuid.New()
	return
}
func (m *TVShow) GetID() uuid.UUID {
	return m.ID
}

func (m *TVShow) GetName() string {
	return m.Name
}
