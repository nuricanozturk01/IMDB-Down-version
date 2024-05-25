package controller

import "imdb_project/service"

type SearchController struct {
	SearchService service.SearchService
}

func NewSearchController(searchService service.SearchService) *SearchController {
	return &SearchController{SearchService: searchService}
}
