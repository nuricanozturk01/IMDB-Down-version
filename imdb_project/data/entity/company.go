package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	MediaID   uuid.UUID
	MediaType string
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}
