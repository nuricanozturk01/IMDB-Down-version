package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rate struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	MediaID   uuid.UUID `gorm:"type:char(36)" json:"media_id"`
	MediaType string
	Rate      float64 `gorm:"type:float;"`
}

func (r *Rate) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
