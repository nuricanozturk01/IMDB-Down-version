package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	MediaID   uuid.UUID
	MediaType string
	URL       string `json:"url"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if photo.ID == uuid.Nil {
		photo.ID = uuid.New()
	}
	return
}
