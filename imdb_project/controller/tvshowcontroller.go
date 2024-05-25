package controller

import "imdb_project/service"

type TVShowController struct {
	TvShowService *service.TvShowService
}

func NewTVShowController(tvShowService *service.TvShowService) *TVShowController {
	return &TVShowController{TvShowService: tvShowService}
}
