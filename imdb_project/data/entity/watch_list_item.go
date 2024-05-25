package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WatchListItem struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	WatchListID uuid.UUID `gorm:"type:char(36)"`
	MediaID     uuid.UUID `gorm:"type:char(36)"`
	MediaType   string
}

func (watchListItem *WatchListItem) BeforeCreate(tx *gorm.DB) (err error) {
	if watchListItem.ID == uuid.Nil {
		watchListItem.ID = uuid.New()
	}
	return
}
