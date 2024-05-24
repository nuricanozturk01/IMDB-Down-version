package entity

import "github.com/google/uuid"

type Rating struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID       uuid.UUID `gorm:"type:char(36)"`
	User         User
	RateableID   uuid.UUID `gorm:"type:char(36)"`
	RateableType string
	Score        int
}
