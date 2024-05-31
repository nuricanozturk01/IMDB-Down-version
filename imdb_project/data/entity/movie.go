package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	ID          uuid.UUID   `json:"id" gorm:"type:char(36);primaryKey;"`
	Name        string      `json:"name" gorm:"type:varchar(80);"`
	Description string      `json:"description" gorm:"type:text;"`
	AverageRate float64     `json:"average_rate" gorm:"type:float;default:0;"`
	Year        uint        `json:"year" gorm:"type:int;"`
	Popularity  uint        `json:"popularity" gorm:"type:int;default:0;"`
	ClickCount  uint        `json:"click_count" gorm:"type:int;default:0;"`
	Trailers    []Trailer   `json:"trailers" gorm:"polymorphic:Media;polymorphicValue:movies"`
	Companies   []Company   `json:"companies" gorm:"polymorphic:Media;polymorphicValue:movies"`
	Celebs      []Celebrity `json:"celebs" gorm:"many2many:movie_celebs;"`
	Photos      []Photo     `json:"photos" gorm:"polymorphic:Media;polymorphicValue:movies"`
	Likes       []Like      `gorm:"polymorphic:Media;polymorphicValue:movies"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}

func (m *Movie) GetID() uuid.UUID {
	return m.ID
}

func (m *Movie) GetName() string {
	return m.Name
}
