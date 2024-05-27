package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"imdb_project/service"
)

type SearchController struct {
	SearchService *service.SearchService
	Validate      *validator.Validate
}

func (c SearchController) SubscribeEndpoints(engine *gin.Engine) {
	engine.GET("/api/v1/search", c.Search)
}

func NewSearchController(searchService *service.SearchService, validator *validator.Validate) *SearchController {
	return &SearchController{SearchService: searchService, Validate: validator}
}

func (c SearchController) Search(context *gin.Context) {
	query := context.Query("keyword")
	response := c.SearchService.Search(query)
	context.JSON(int(response.StatusCode), response)
}
