package tests

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	helper "imdb_project/data/dal"
	"imdb_project/data/entity"
	"imdb_project/service"
	"log"
	"os"
)

type AppContext struct {
	DB                    *gorm.DB
	MovieService          *service.MovieService
	TvShowService         *service.TvShowService
	SearchService         *service.SearchService
	UserService           *service.UserService
	AuthenticationService *service.AuthenticationService
	CelebrityService      *service.CelebrityService
}

func configure() *AppContext {
	db, err := initTestDB()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Service Helper (Facade Pattern) (for Repository Layer)
	imdbHelper := helper.New(db)

	// Service Layer
	movieService := service.NewMovieService(imdbHelper, nil)
	tvShowService := service.NewTvShowService(imdbHelper, nil)
	searchService := service.NewSearchService(imdbHelper)
	userService := service.NewUserService(imdbHelper)
	authenticationService := service.NewAuthenticationService(imdbHelper)
	celebrityService := service.NewCelebrityService(imdbHelper)

	return &AppContext{
		DB:                    db,
		MovieService:          movieService,
		TvShowService:         tvShowService,
		SearchService:         searchService,
		UserService:           userService,
		AuthenticationService: authenticationService,
		CelebrityService:      celebrityService,
	}
}

func initTestDB() (*gorm.DB, error) {
	databaseConnection := os.Getenv("TEST_DB_DSN")

	var err error

	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Rate{},
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Like{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.WatchList{},
		&entity.WatchListItem{},
	)

	if err != nil {
		log.Println("Failed to migrate database:", err)
		return nil, err
	}

	log.Println("Database migrated successfully")

	return db, nil
}
func resetTestDB() error {
	databaseConnection := os.Getenv("TEST_DB_DSN")

	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop all tables
	err = db.Migrator().DropTable(
		&entity.User{},
		&entity.Rate{},
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Like{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.WatchList{},
		&entity.WatchListItem{},
	)

	log.Println("Database tables dropped successfully")

	// Recreate all tables
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Rate{},
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Like{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.WatchList{},
		&entity.WatchListItem{},
	)
	if err != nil {
		return err
	}

	log.Println("Database migrated successfully")

	return nil
}
