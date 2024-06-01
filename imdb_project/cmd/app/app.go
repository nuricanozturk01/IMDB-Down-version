package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jasonlvhit/gocron"
	"imdb_project/config"
	databaseConfig "imdb_project/config/database"
	"imdb_project/config/middleware"
	"imdb_project/config/sqs"
	"imdb_project/controller"
	helper "imdb_project/data/dal"
	"imdb_project/data/entity/enum"
	"imdb_project/service"
	"log"
	"os"
	"strconv"
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

	// CORS configuration
	engine.Use(middleware.CorsPolicy())

	// Start the session store
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	// Service Helper (Facade Pattern) (for Repository Layer)
	imdbHelper := helper.New(db)

	// Service Layer
	queueService := sqs.NewSQSClient(os.Getenv("QUEUE_URL"))
	movieService := service.NewMovieService(imdbHelper, queueService)
	tvShowService := service.NewTvShowService(imdbHelper, queueService)
	searchService := service.NewSearchService(imdbHelper)
	userService := service.NewUserService(imdbHelper)
	authenticationService := service.NewAuthenticationService(imdbHelper)
	celebrityService := service.NewCelebrityService(imdbHelper)

	// Middleware Layer and cors settings
	authMiddleware := middleware.NewAuthMiddleware(store)

	// Controller Layer
	authController := controller.NewAuthController(authenticationService, validate, store)
	userController := controller.NewUserController(userService, validate)
	movieController := controller.NewMovieController(movieService, validate, store)
	tvShowController := controller.NewTVShowController(tvShowService, validate, store)
	searchController := controller.NewSearchController(searchService, celebrityService, validate, store)

	// Public routes
	authController.SubscribeEndpoints(engine)
	engine.Use(authMiddleware.Middleware())

	// Protected routes
	protected := engine.Group("/")
	// Protected routes subscription (session logic)
	userController.SubscribeEndpoints(protected)
	movieController.SubscribeEndpoints(protected)
	tvShowController.SubscribeEndpoints(protected)
	searchController.SubscribeEndpoints(protected)

	go startScheduler(queueService, imdbHelper)

	// Run the server
	err = engine.Run(":5050")

	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}

}

func startScheduler(queueService *sqs.QueueService, imdbHelper *helper.ImdbHelper) {
	scheduler := gocron.NewScheduler()

	defer func() { scheduler.Clear() }()

	timeMinute, _ := strconv.ParseUint(os.Getenv("QUEUE_TIME"), 10, 64)

	err := scheduler.Every(timeMinute).Minute().Do(processMessages, queueService, imdbHelper)

	if err != nil {
		fmt.Printf("Problem occurred while creating job:%s\n", err.Error())
		os.Exit(1)
	}
	scheduler.Start()

	select {}
}

func processMessages(sqsClient *sqs.QueueService, imdbHelper *helper.ImdbHelper) {

	ctx := context.Background()

	messages, err := sqsClient.ReceiveMessages(ctx)

	if err != nil {
		log.Println("Error receiving messages:", err)
		return
	}

	for _, message := range messages {
		var msg service.GenericMessage

		_ = json.Unmarshal([]byte(*message.Body), &msg) // error is unexpected, so I don't handle it

		switch msg.Type {
		case enum.MovieType:
			updateMoviePopularity(imdbHelper, msg)
		case enum.TvShowType:
			updateTvShowPopularity(imdbHelper, msg)
		default:
			log.Println("Unknown message type")
		}

		if err := sqsClient.DeleteMessage(ctx, message.ReceiptHandle); err != nil {
			log.Println("Error deleting message:", err)
		}

		log.Println("Message processed successfully")
	}
}

func updateTvShowPopularity(imdbHelper *helper.ImdbHelper, msg service.GenericMessage) {
	tvShow := imdbHelper.FindTvShowByID(uuid.MustParse(msg.Id))
	tvShow.Popularity++
	_, err := imdbHelper.TvShowRepository.Update(tvShow)
	if err != nil {
		log.Println("Error saving tv show to DB:", err)
	}
}

func updateMoviePopularity(imdbHelper *helper.ImdbHelper, msg service.GenericMessage) {
	movie := imdbHelper.FindMovieByID(uuid.MustParse(msg.Id))

	movie.Popularity++
	_, err := imdbHelper.MovieRepository.Update(movie)

	if err != nil {
		log.Println("Error saving movie to DB:", err)
	}

}
