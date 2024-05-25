package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36)"`
	MediaID   uuid.UUID `gorm:"type:char(36)"`
	MediaType string
}

func (l *Like) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
