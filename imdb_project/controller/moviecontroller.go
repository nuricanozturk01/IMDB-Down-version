package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"imdb_project/data/dto"
	"imdb_project/service"
	"net/http"
	"strconv"
)

type MovieController struct {
	MovieService service.IMovieService
	Validate     *validator.Validate
	Store        *sessions.CookieStore
}

func (c *MovieController) SubscribeEndpoints(engine *gin.RouterGroup) {
	engine.POST("/api/v1/movie/create", c.CreateMovie)
	engine.POST("/api/v1/movie/like", c.LikeMovie)
	engine.POST("/api/v1/movie/dislike", c.DislikeMovie)
	engine.GET("/api/v1/movie/all", c.FindAllMovies)
	engine.GET("/api/v1/movie", c.FindMovieById)
	engine.POST("/api/v1/movie/watchlist/add", c.AddMovieToWatchList)
	engine.POST("/api/v1/movie/rate", c.RateMovie)
	engine.DELETE("/api/v1/movie/watchlist/delete", c.RemoveMovieFromWatchList)
}

func NewMovieController(movieService service.IMovieService, validator *validator.Validate, store *sessions.CookieStore) *MovieController {
	return &MovieController{MovieService: movieService, Validate: validator, Store: store}
}

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie *dto.MovieCreateDTO

	err := ctx.BindJSON(&movie)

	if validationErr := c.Validate.Struct(movie); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.MovieService.CreateMovie(movie)

	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) LikeMovie(ctx *gin.Context) {

	userID := c.getUserID(ctx)
	movieId, _ := uuid.Parse(ctx.Query("movie_id"))

	response := c.MovieService.LikeMovie(movieId, userID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) DislikeMovie(ctx *gin.Context) {

	userID := c.getUserID(ctx)
	movieId, _ := uuid.Parse(ctx.Query("movie_id"))

	response := c.MovieService.DislikeMovie(movieId, userID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) FindAllMovies(ctx *gin.Context) {
	response := c.MovieService.FindAllMovies()
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) FindMovieById(ctx *gin.Context) {
	id := ctx.Query("id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}
	response := c.MovieService.FindMovieById(parsedUUID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) AddMovieToWatchList(ctx *gin.Context) {
	session, err := c.Store.Get(ctx.Request, "imdb-session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session: " + err.Error()})
		return
	}

	userID, ok := session.Values["id"]
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in session"})
		return
	}

	movieId, _ := uuid.Parse(ctx.Query("movie_id"))
	userIdUUID := uuid.MustParse(userID.(string))

	response := c.MovieService.AddMovieToWatchList(userIdUUID, movieId)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) RemoveMovieFromWatchList(ctx *gin.Context) {
	userId := c.getUserID(ctx)
	movieId, _ := uuid.Parse(ctx.Query("movie_id"))

	response := c.MovieService.RemoveMovieFromWatchList(movieId, userId)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) RateMovie(context *gin.Context) {
	userID := c.getUserID(context)
	movieID, _ := uuid.Parse(context.Query("movie_id"))
	rate, _ := context.GetQuery("rate")

	rateFloat, err := strconv.ParseFloat(rate, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.MovieService.RateMovie(movieID, userID, rateFloat)
	context.JSON(int(response.StatusCode), response)
}

//...

func (c *MovieController) getUserID(ctx *gin.Context) uuid.UUID {
	session, err := c.Store.Get(ctx.Request, "imdb-session")
	if err != nil {
		return uuid.UUID{}
	}

	userID, _ := session.Values["id"]

	return uuid.MustParse(userID.(string))
}
