package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trailer struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	MediaID   uuid.UUID
	MediaType string
	URL       string `json:"url"`
}

func (trailer *Trailer) BeforeCreate(tx *gorm.DB) (err error) {
	trailer.ID = uuid.New()
	return
}
