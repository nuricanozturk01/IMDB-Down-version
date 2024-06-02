package entity

type Country struct {
	CountryCode string `gorm:"primaryKey;type:varchar(2)"`
	CountryName string `gorm:"type:varchar(100);not null"`
	Cities      []City `gorm:"foreignKey:CountryCode"`
}
