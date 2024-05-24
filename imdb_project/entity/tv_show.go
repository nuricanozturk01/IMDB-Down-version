package entity

/*type TvShow struct {
	gorm.Model
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
	Celebs         []Celebrity     `json:"celebs" gorm:"many2many:movie_celebs;"`
	WatchListItems []WatchListItem `json:"watch_list_items" gorm:"foreignKey:ItemID"`
	Ratings        []Rating        `json:"ratings" gorm:"foreignKey:RateableID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
*/
