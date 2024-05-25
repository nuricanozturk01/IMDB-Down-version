package entity

import (
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	MediaID   uuid.UUID
	MediaType string
}
