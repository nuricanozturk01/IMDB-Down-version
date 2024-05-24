package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID       `json:"id" gorm:"primaryKey;type:char(36);"`
	Username       string          `json:"username" gorm:"unique;not null;varchar(45)"`
	Password       string          `json:"password" gorm:"not null;"`
	FirstName      string          `json:"first_name" gorm:"not null;varchar(45)"`
	LastName       string          `json:"last_name" gorm:"not null;varchar(45)"`
	Email          string          `json:"email" gorm:"not null;varchar(80)"`
	WatchList      Watchlist       `json:"watch_list"`
	WatchListItems []WatchListItem `json:"watch_list_items"`
	Ratings        []Rating        `json:"ratings" gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Movie struct {
	ID             uuid.UUID       `json:"id" gorm:"type:char(36);primaryKey;"`
	Name           string          `json:"name" gorm:"type:varchar(80);"`
	AverageRate    float64         `json:"average_rate"`
	Year           uint            `json:"year" gorm:"type:int;"`
	Popularity     uint            `json:"popularity" gorm:"type:int;"`
	ClickCount     uint            `json:"click_count" gorm:"type:int;"`
	Trailers       []Trailer       `json:"trailers" gorm:"polymorphic:Owner;"`
	Celebs         []Celebrity     `json:"celebs" gorm:"many2many:movie_celebs;"`
	WatchListItems []WatchListItem `json:"watch_list_items" gorm:"foreignKey:ItemID"`
	Ratings        []Rating        `json:"ratings" gorm:"foreignKey:RateableID"`
	Photos         []Photo         `json:"photos" gorm:"polymorphic:Owner"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TVShow struct {
	ID             uuid.UUID       `json:"id" gorm:"type:char(36);primaryKey;"`
	Name           string          `json:"name" gorm:"type:varchar(80);"`
	Year           int             `json:"year" gorm:"type:int;"`
	Popularity     int             `json:"popularity" gorm:"type:int;"`
	AverageRate    float64         `json:"average_rate"`
	ClickCount     uint32          `json:"click_count" gorm:"type:int;"`
	EpisodeCount   uint32          `json:"episode_count" gorm:"type:int;"`
	SeasonCount    uint32          `json:"season_count" gorm:"type:int;"`
	Photos         []Photo         `json:"photos" gorm:"polymorphic:Owner"`
	Trailers       []Trailer       `json:"trailers" gorm:"polymorphic:Owner;"`
	Celebs         []Celebrity     `json:"celebs" gorm:"many2many:tvshow_celebs;"`
	WatchListItems []WatchListItem `json:"watch_list_items" gorm:"foreignKey:ItemID"`
	Ratings        []Rating        `json:"ratings" gorm:"foreignKey:RateableID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Celebrity struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key"`
	Name    string
	Movies  []Movie  `gorm:"many2many:movie_celebs;"`
	TVShows []TVShow `gorm:"many2many:tvshow_celebs;"`
}

type Company struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"unique"`
}

type Photo struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	URL       string `json:"url"`
}

type Rating struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID       uuid.UUID `gorm:"type:char(36)"`
	User         User
	RateableID   uuid.UUID `gorm:"type:char(36)"`
	RateableType string
	Score        int
}

type Trailer struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	URL       string
}

type WatchListItem struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	OwnerID   uuid.UUID `gorm:"type:char(36)"`
	OwnerType string
	ItemID    uuid.UUID
	ItemType  string
	Title     string
	UserID    uuid.UUID `gorm:"type:char(36)"`
	User      User
	Ratings   []Rating `gorm:"polymorphic:Rateable;"`
}

type Watchlist struct {
	ID     uuid.UUID       `gorm:"type:char(36);primary_key"`
	UserID uuid.UUID       `gorm:"type:char(36)"`
	Items  []WatchListItem `gorm:"polymorphic:Owner;"`
}
