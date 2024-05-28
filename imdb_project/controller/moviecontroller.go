package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"imdb_project/data/dto"
	"imdb_project/service"
	"net/http"
)

type MovieController struct {
	MovieService service.IMovieService
	Validate     *validator.Validate
}

func (c *MovieController) SubscribeEndpoints(engine *gin.RouterGroup) {
	engine.POST("/api/v1/movie/create", c.CreateMovie)
	engine.POST("/api/v1/movie/like", c.LikeMovie)
	engine.POST("/api/v1/movie/dislike", c.DislikeMovie)
	engine.GET("/api/v1/movie/all", c.FindAllMovies)
	engine.GET("/api/v1/movie/:id", c.FindMovieById)
	engine.POST("/api/v1/movie/watchlist/add", c.AddMovieToWatchList)
	engine.DELETE("/api/v1/movie/watchlist/delete", c.RemoveMovieFromWatchList)
}

func NewMovieController(movieService service.IMovieService, validator *validator.Validate) *MovieController {
	return &MovieController{MovieService: movieService, Validate: validator}
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
	var like *dto.LikeDTO
	err := ctx.BindJSON(&like)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := c.MovieService.LikeMovie(like.ID, like.UserID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) DislikeMovie(ctx *gin.Context) {
	var dislike *dto.LikeDTO
	err := ctx.BindJSON(&dislike)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := c.MovieService.DislikeMovie(dislike.ID, dislike.UserID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) FindAllMovies(ctx *gin.Context) {
	response := c.MovieService.FindAllMovies()
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) FindMovieById(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}
	response := c.MovieService.FindMovieById(parsedUUID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) AddMovieToWatchList(ctx *gin.Context) {
	var watchList *dto.WatchListItemIoDTO
	err := ctx.BindJSON(&watchList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := c.MovieService.AddMovieToWatchList(watchList.UserID, watchList.MediaID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *MovieController) RemoveMovieFromWatchList(ctx *gin.Context) {
	var watchList *dto.WatchListItemIoDTO
	err := ctx.BindJSON(&watchList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := c.MovieService.RemoveMovieFromWatchList(watchList.UserID, watchList.MediaID)
	ctx.JSON(int(response.StatusCode), response)
}

//...
