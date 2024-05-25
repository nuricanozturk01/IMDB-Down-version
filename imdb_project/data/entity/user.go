package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:char(36);"`
	Username  string    `json:"username" gorm:"unique;not null;varchar(45)"`
	Password  string    `json:"password" gorm:"not null;"`
	FirstName string    `json:"first_name" gorm:"not null;varchar(45)"`
	LastName  string    `json:"last_name" gorm:"not null;varchar(45)"`
	Email     string    `json:"email" gorm:"not null;varchar(80)"`
	WatchList WatchList `gorm:"foreignKey:UserID"`
	Likes     []Like    `gorm:"foreignKey:UserID"`
	Photos    []Photo   `json:"photos" gorm:"polymorphic:Media;polymorphicValue:users"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate Like Trigger in SQL
func (usr *User) BeforeCreate(tx *gorm.DB) (err error) {
	usr.ID = uuid.New()
	usr.WatchList = WatchList{ID: uuid.New()}
	return
}
func (usr *User) GetID() uuid.UUID {
	return usr.ID
}

func (usr *User) GetName() string {
	return usr.Username
}
