package entity

import "github.com/google/uuid"

type Company struct {
	ID      uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name    string    `json:"name" gorm:"unique"`
	OwnerID uuid.UUID `gorm:"type:char(36)"` // OwnerID is the ID of the movie or tv show
}
