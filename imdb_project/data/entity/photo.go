package entity

import "github.com/google/uuid"

type Photo struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	URL       string `json:"url"`
}
