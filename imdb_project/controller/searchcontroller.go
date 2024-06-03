package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"imdb_project/service"
)

type SearchController struct {
	SearchService    *service.SearchService
	CelebrityService *service.CelebrityService
	Validate         *validator.Validate
	Store            *sessions.CookieStore
}

func (c SearchController) SubscribeEndpoints(engine *gin.RouterGroup, protected *gin.RouterGroup) {
	// public
	engine.GET("/search", c.Search)
	engine.GET("/celebrity", c.FindCelebrityByID)
	// protected
	protected.GET("/celebrity/all", c.FindAllCelebrities)
	protected.GET("/watchlist", c.FindWatchList)
}

func NewSearchController(searchService *service.SearchService, celebrityService *service.CelebrityService, validator *validator.Validate,
	Store *sessions.CookieStore) *SearchController {
	return &SearchController{SearchService: searchService, CelebrityService: celebrityService, Validate: validator, Store: Store}
}

func (c SearchController) Search(context *gin.Context) {
	query := context.Query("keyword")
	response := c.SearchService.Search(query)
	context.JSON(int(response.StatusCode), response)
}

func (c SearchController) FindAllCelebrities(context *gin.Context) {
	response := c.CelebrityService.FindAllCelebrities()
	context.JSON(int(response.StatusCode), response)
}

func (c SearchController) FindWatchList(context *gin.Context) {
	userID := c.getUserID(context)
	response := c.SearchService.FindWatchList(userID)
	context.JSON(int(response.StatusCode), response)
}

func (c SearchController) getUserID(ctx *gin.Context) uuid.UUID {
	session, err := c.Store.Get(ctx.Request, "imdb-session")
	if err != nil {
		return uuid.UUID{}
	}

	userID, _ := session.Values["id"]

	return uuid.MustParse(userID.(string))
}

func (c SearchController) FindCelebrityByID(context *gin.Context) {
	id := context.Query("id")
	fmt.Println(id)
	response := c.CelebrityService.FindCelebrityByID(uuid.MustParse(id))
	context.JSON(int(response.StatusCode), response)
}
