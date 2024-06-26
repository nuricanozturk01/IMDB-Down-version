package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"imdb_project/data/dto"
	"imdb_project/service"
	"net/http"
	"strconv"
)

type TVShowController struct {
	TvShowService *service.TvShowService
	Validate      *validator.Validate
	Store         *sessions.CookieStore
}

func NewTVShowController(tvShowService *service.TvShowService, validator *validator.Validate, store *sessions.CookieStore) *TVShowController {
	return &TVShowController{TvShowService: tvShowService, Validate: validator, Store: store}
}

func (c *TVShowController) SubscribeEndpoints(engine *gin.RouterGroup, protected *gin.RouterGroup) {

	// public
	engine.GET("/tv_show/all", c.FindAllTvShows)
	engine.GET("/tv_show", c.FindTvShowById)

	// protected
	protected.POST("/tv_show/create", c.CreateTvShow)
	protected.POST("/tv_show/like", c.LikeTvShow)
	protected.POST("/tv_show/dislike", c.DislikeTvShow)
	protected.POST("/tv_show/watchlist", c.AddTvShowToWatchList)
	protected.POST("/tv_show/rate", c.RateTvShow)
	protected.DELETE("/tv_show/watchlist", c.RemoveTvShowFromWatchList)
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
	id := uuid.MustParse(context.Query("id"))

	response := c.TvShowService.FindTvShowById(id)
	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) AddTvShowToWatchList(context *gin.Context) {

	session, err := c.Store.Get(context.Request, "imdb-session")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session: " + err.Error()})
		return
	}

	userID, ok := session.Values["id"]
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in session"})
		return
	}

	tvShowId, _ := uuid.Parse(context.Query("tv_show_id"))
	userIdUUID := uuid.MustParse(userID.(string))

	response := c.TvShowService.AddTvShowToWatchList(userIdUUID, tvShowId)

	context.JSON(int(response.StatusCode), response)

}

func (c *TVShowController) RemoveTvShowFromWatchList(context *gin.Context) {

	session, err := c.Store.Get(context.Request, "imdb-session")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session: " + err.Error()})
		return
	}

	userID, ok := session.Values["id"]
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in session"})
		return
	}

	tvShowId, _ := uuid.Parse(context.Query("tv_show_id"))

	response := c.TvShowService.RemoveTvShowFromWatchList(uuid.MustParse(userID.(string)), tvShowId)

	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) RateTvShow(context *gin.Context) {

	userID := c.getUserID(context)
	tvShowId, _ := uuid.Parse(context.Query("tv_show_id"))
	rate, _ := context.GetQuery("rate")

	rateFloat, err := strconv.ParseFloat(rate, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.TvShowService.RateTvShow(tvShowId, userID, rateFloat)

	context.JSON(int(response.StatusCode), response)
}

func (c *TVShowController) getUserID(ctx *gin.Context) uuid.UUID {
	session, err := c.Store.Get(ctx.Request, "imdb-session")
	if err != nil {
		return uuid.UUID{}
	}

	userID, _ := session.Values["id"]

	return uuid.MustParse(userID.(string))
}
