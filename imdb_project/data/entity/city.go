package entity

type City struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	CityName    string `gorm:"type:varchar(200);not null"`
	CountryCode string `gorm:"type:varchar(2);not null"`
}
