package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WatchList struct {
	ID     uuid.UUID       `gorm:"type:char(36);primaryKey"`
	UserID uuid.UUID       `gorm:"type:char(36)"`
	Items  []WatchListItem `gorm:"foreignKey:WatchListID"`
}

func (watchList *WatchList) BeforeCreate(tx *gorm.DB) (err error) {
	if watchList.ID == uuid.Nil {
		watchList.ID = uuid.New()
	}
	return
}
