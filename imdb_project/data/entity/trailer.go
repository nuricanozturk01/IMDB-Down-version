package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trailer struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	MediaID   uuid.UUID `json:"media_id"`
	MediaType string    `json:"media_type"`
	URL       string    `json:"url"`
}

func (trailer *Trailer) BeforeCreate(tx *gorm.DB) (err error) {
	if trailer.ID == uuid.Nil {
		trailer.ID = uuid.New()
	}
	return
}
