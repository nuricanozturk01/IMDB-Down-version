package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
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
	"time"
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

	// Cache Layer
	cache, err := bigcache.New(context.Background(), bigcache.Config{
		Shards:             1024,
		LifeWindow:         24 * time.Hour,
		CleanWindow:        1 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		HardMaxCacheSize:   2048,
	})
	if err != nil {
		log.Fatal(err)
	}
	// CORS configuration
	engine.Use(middleware.CorsPolicy())

	// Start the session store
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	// Service Helper (Facade Pattern) (for Repository Layer)
	imdbHelper := helper.New(db, cache)

	// Service Layer
	queueService := sqs.NewSQSClient(os.Getenv("QUEUE_URL"))
	movieService := service.NewMovieService(imdbHelper, queueService)
	tvShowService := service.NewTvShowService(imdbHelper, queueService)
	searchService := service.NewSearchService(imdbHelper)
	userService := service.NewUserService(imdbHelper)
	authenticationService := service.NewAuthenticationService(imdbHelper)
	celebrityService := service.NewCelebrityService(imdbHelper)
	informationService := service.NewInformationService(imdbHelper, cache)

	// Middleware Layer and cors settings
	//authMiddleware := middleware.NewAuthMiddleware(store)

	// Public routes group
	public := engine.Group("/api/v1/public")

	// Private routes group
	private := engine.Group("/api/v1/private")
	//private.Use(authMiddleware.Middleware())

	// Controller Layer
	authController := controller.NewAuthController(authenticationService, informationService, validate, store)
	userController := controller.NewUserController(userService, validate, store)
	movieController := controller.NewMovieController(movieService, validate, store)
	tvShowController := controller.NewTVShowController(tvShowService, validate, store)
	searchController := controller.NewSearchController(searchService, celebrityService, validate, store)

	// Public routes
	authController.SubscribeEndpoints(public)

	// Private routes subscription (session logic)
	userController.SubscribeEndpoints(private)
	movieController.SubscribeEndpoints(public, private)
	tvShowController.SubscribeEndpoints(public, private)
	searchController.SubscribeEndpoints(public, private)

	go startScheduler(queueService, imdbHelper)
	go cacheAllCountries(cache, imdbHelper)

	// Run the server
	err = engine.Run(":5050")

	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}

}

func cacheAllCountries(cache *bigcache.BigCache, imdbHelper *helper.ImdbHelper) {
	countries, err := imdbHelper.CountryRepository.FindAll()
	if err != nil {
		log.Println("Error while fetching countries:", err)
		return
	}

	data, err := json.Marshal(countries)
	if err != nil {
		log.Println("Error while marshalling countries:", err)
		return
	}

	err = cache.Set("countries", data)
	if err != nil {
		log.Println("Error while setting cache:", err)
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
