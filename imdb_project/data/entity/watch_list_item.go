package entity

import "github.com/google/uuid"

type WatchListItem struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	ItemID    uuid.UUID
	ItemType  string
	Title     string
	UserID    uuid.UUID `gorm:"type:char(36)"`
	User      User
	Ratings   []Rating `gorm:"polymorphic:Rateable;"`
}
