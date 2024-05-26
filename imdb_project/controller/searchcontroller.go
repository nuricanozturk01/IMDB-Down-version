package controller

import (
	"github.com/gin-gonic/gin"
	"imdb_project/service"
)

type SearchController struct {
	SearchService *service.SearchService
}

func (c SearchController) SubscribeEndpoints(engine *gin.Engine) {
	engine.GET("/search", c.Search)
}

func NewSearchController(searchService *service.SearchService) *SearchController {
	return &SearchController{SearchService: searchService}
}

func (c SearchController) Search(context *gin.Context) {
	query := context.Query("query")
	response := c.SearchService.Search(query)
	context.JSON(int(response.StatusCode), response)
}
