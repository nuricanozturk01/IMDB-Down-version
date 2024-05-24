package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imdb_project/data/entity"
	"log"
	"os"
)

type ImdbHelper struct {
	db *gorm.DB
}

func InitDb() (*gorm.DB, error) {
	databaseConnection := os.Getenv("DB_DSN")

	var err error

	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Rating{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.User{},
		&entity.Watchlist{},
		&entity.WatchListItem{},
	)

	if err != nil {
		log.Println("Failed to migrate database:", err)
		return nil, err
	}

	log.Println("Database migrated successfully")

	return db, nil
}
