package entity

import "github.com/google/uuid"

type Trailer struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	URL       string `json:"url"`
}
