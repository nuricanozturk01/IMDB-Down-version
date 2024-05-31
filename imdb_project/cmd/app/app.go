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
	engine := gin.Default()

	// CORS configuration
	engine.Use(middleware.CorsPolicy())

	// Start the session store
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

	// Middleware Layer and cors settings
	authMiddleware := middleware.NewAuthMiddleware(store)

	// Controller Layer
	authController := controller.NewAuthController(authenticationService, validate, store)
	userController := controller.NewUserController(userService, validate)
	movieController := controller.NewMovieController(movieService, validate)
	tvShowController := controller.NewTVShowController(tvShowService, validate)
	searchController := controller.NewSearchController(searchService, celebrityService, validate)

	// Public routes
	authController.SubscribeEndpoints(engine)
	engine.Use(authMiddleware.Middleware())

	// Protected routes
	protected := engine.Group("/")
	//protected.Use(authMiddleware.Middleware())

	userController.SubscribeEndpoints(protected)
	movieController.SubscribeEndpoints(protected)
	tvShowController.SubscribeEndpoints(protected)
	searchController.SubscribeEndpoints(protected)

	// Run the server
	err = engine.Run(":5050")

	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}
}
