package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"imdb_project/config"
	databaseConfig "imdb_project/config/database"
	"imdb_project/config/middleware"
	"imdb_project/controller"
	helper "imdb_project/data/dal"
	"imdb_project/service"
	"log"
	"os"
)

var validate *validator.Validate
var store *sessions.CookieStore

func init() {
	validate = validator.New()
}

func Run() {
	// Load the configuration
	config.Load()

	// Initialize the database
	db, err := databaseConfig.InitDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// HTTP Layer
	engine := gin.New()

	// start the session store
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	// Service Helper (Facade Pattern) (for Repository Layer)
	imdbHelper := helper.New(db)

	// Service Layer
	movieService := service.NewMovieService(imdbHelper)
	tvShowService := service.NewTvShowService(imdbHelper)
	searchService := service.NewSearchService(imdbHelper)
	userService := service.NewUserService(imdbHelper)
	authenticationService := service.NewAuthenticationService(imdbHelper)
	celebrityService := service.NewCelebrityService(imdbHelper)

	// Middleware Layer
	authMiddleware := middleware.NewAuthMiddleware(store)

	// Controller Layer
	authController := controller.NewAuthController(authenticationService, validate, store)
	userController := controller.NewUserController(userService, validate)
	movieController := controller.NewMovieController(movieService, validate)
	tvShowController := controller.NewTVShowController(tvShowService, validate)
	searchController := controller.NewSearchController(searchService, celebrityService, validate)

	// Public routes
	authController.SubscribeEndpoints(engine)

	// Protected routes
	protected := engine.Group("/")
	protected.Use(authMiddleware.Middleware())

	userController.SubscribeEndpoints(protected)
	movieController.SubscribeEndpoints(protected)
	tvShowController.SubscribeEndpoints(protected)
	searchController.SubscribeEndpoints(protected)

	// Run the server
	err = engine.Run(":5050")

	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}

	/*movie, err := imdbHelper.FindAllMovies()
	if err != nil {
		log.Fatal("Failed to fetch movies:", err)
	}

	firstMovie := movie[0]

	celeb, err := imdbHelper.FindAllCelebrities()

	if err != nil {
		log.Fatal("Failed to fetch celebrities:", err)
	}

	firstCeleb, err := imdbHelper.FindCelebrityByID(celeb[0].ID)

	if err != nil {
		log.Fatal("Failed to fetch celebrity:", err)

	}

	fmt.Println(firstMovie.Name)
	fmt.Println(firstCeleb.ID)

	err = db.Model(&firstMovie).Association("Celebs").Replace(firstCeleb)
	if err != nil {
		log.Fatal("Failed to associate movie with celebrity:", err)
	}*/

	/*imdbHelper := helper.New(db)

	result := imdbHelper.Search("Inc")

	log.Println(result)

	movie := entity.Movie{
		Name: "Inception",
		Photos: []entity.Photo{
			{URL: "photo1.jpg"},
			{URL: "photo2.jpg"},
		},
		Trailers: []entity.Trailer{
			{URL: "trailer1.mp4"},
			{URL: "trailer2.mp4"},
		},
	}

	tvShow := entity.TVShow{
		Name: "Breaking Bad",
		Photos: []entity.Photo{
			{URL: "photo1.jpg"},
			{URL: "photo2.jpg"},
		},
		Trailers: []entity.Trailer{
			{URL: "trailer1.mp4"},
			{URL: "trailer2.mp4"},
		},
	}

	user := entity.User{
		Username: "John Doe",
		Photos: []entity.Photo{
			{URL: "user_photo1.jpg"},
			{URL: "user_photo2.jpg"},
		},
	}

	celebrity := entity.Celebrity{
		Name: "Jane Smith",
		Photos: []entity.Photo{
			{URL: "celeb_photo1.jpg"},
			{URL: "celeb_photo2.jpg"},
		},
	}

	db.Create(&movie)
	db.Create(&tvShow)
	db.Create(&user)
	db.Create(&celebrity)

	watchListItem1 := entity.WatchListItem{
		WatchListID: user.WatchList.ID,
		MediaID:     movie.ID,
		MediaType:   "movies",
	}
	watchListItem2 := entity.WatchListItem{
		WatchListID: user.WatchList.ID,
		MediaID:     tvShow.ID,
		MediaType:   "tv_shows",
	}

	db.Create(&watchListItem1)
	db.Create(&watchListItem2)

	like1 := entity.Like{
		UserID:    user.ID,
		MediaID:   movie.ID,
		MediaType: "movies",
	}
	db.Create(&like1)

	like2 := entity.Like{
		UserID:    user.ID,
		MediaID:   tvShow.ID,
		MediaType: "tv_shows",
	}
	db.Create(&like2)*/
}
