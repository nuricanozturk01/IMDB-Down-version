package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:char(36);"`
	Username  string    `json:"username" gorm:"unique;not null;varchar(45)"`
	Password  string    `json:"password" gorm:"not null;"`
	FirstName string    `json:"first_name" gorm:"not null;varchar(45)"`
	LastName  string    `json:"last_name" gorm:"not null;varchar(45)"`
	Email     string    `json:"email" gorm:"not null;varchar(80)"`
	WatchList Watchlist `json:"watch_list"`
	Ratings   []Rating  `json:"ratings" gorm:"foreignKey:UserID"`
	Photos    []Photo   `json:"photos" gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
