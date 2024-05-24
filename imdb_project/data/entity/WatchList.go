package entity

import "github.com/google/uuid"

type Watchlist struct {
	ID     uuid.UUID       `gorm:"type:char(36);primaryKey"`
	UserID uuid.UUID       `gorm:"type:char(36)"`
	Items  []WatchListItem `gorm:"polymorphic:Owner;"`
}
