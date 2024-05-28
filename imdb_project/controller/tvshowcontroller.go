package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"imdb_project/data/dto"
	"imdb_project/service"
)

type TVShowController struct {
	TvShowService *service.TvShowService
	Validate      *validator.Validate
}

func NewTVShowController(tvShowService *service.TvShowService, validator *validator.Validate) *TVShowController {
	return &TVShowController{TvShowService: tvShowService, Validate: validator}
}

func (c *TVShowController) SubscribeEndpoints(engine *gin.RouterGroup) {
	engine.POST("/api/v1/tv_show/create", c.CreateTvShow)
	engine.POST("/api/v1/tv_show/like", c.LikeTvShow)
	engine.POST("/api/v1/tv_show/dislike", c.DislikeTvShow)
	engine.GET("/api/v1/tv_show/all", c.FindAllTvShows)
	engine.GET("/api/v1/tv_show/:id", c.FindTvShowById)
	engine.POST("/api/v1/tv_show/watchlist", c.AddTvShowToWatchList)
	engine.DELETE("/api/v1/tv_show/watchlist", c.RemoveTvShowFromWatchList)
}

func (c *TVShowController) CreateTvShow(context *gin.Context) {
	var tvShow *dto.TvShowCreateDTO

	err := context.BindJSON(&tvShow)

	if validationErr := c.Validate.Struct(tvShow); validationErr != nil {
		context.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.CreateTvShow(tvShow)

	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) LikeTvShow(context *gin.Context) {
	var like *dto.LikeDTO

	err := context.BindJSON(&like)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.LikeTvShow(like.ID, like.UserID)

	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) DislikeTvShow(context *gin.Context) {
	var dislike *dto.LikeDTO

	err := context.BindJSON(&dislike)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.DislikeTvShow(dislike.ID, dislike.UserID)

	context.JSON(int(response.StatusCode), response)

}

func (c *TVShowController) FindAllTvShows(context *gin.Context) {

	response := c.TvShowService.FindAllTvShow()

	context.JSON(int(response.StatusCode), response)

}

func (c *TVShowController) FindTvShowById(context *gin.Context) {

	id := context.Param("id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}
	response := c.TvShowService.FindTvShowById(parsedUUID)
	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) AddTvShowToWatchList(context *gin.Context) {

	var watchList *dto.WatchListItemIoDTO

	err := context.BindJSON(&watchList)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.AddTvShowToWatchList(watchList.UserID, watchList.MediaID)

	context.JSON(int(response.StatusCode), response)

}

func (c *TVShowController) RemoveTvShowFromWatchList(context *gin.Context) {

	var watchList *dto.WatchListItemIoDTO

	err := context.BindJSON(&watchList)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.RemoveTvShowFromWatchList(watchList.UserID, watchList.MediaID)

	context.JSON(int(response.StatusCode), response)

}
