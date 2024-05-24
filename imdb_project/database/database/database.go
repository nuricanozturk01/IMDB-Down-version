package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"imdb_project/entity"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	databaseConnection := os.Getenv("DB_DSN")

	var err error

	DB, err = gorm.Open(mysql.Open(databaseConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Rating{},
		&entity.Trailer{},
		&entity.TvShow{},
		&entity.User{},
		&entity.Watchlist{},
	)

	if err != nil {
		log.Println("Failed to migrate database:", err)
		return
	}

	log.Println("Database migrated successfully")
}
