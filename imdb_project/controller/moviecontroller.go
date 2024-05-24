package controller

import "imdb_project/service"

type MovieController struct {
	MovieService service.MovieService
}

func New() *MovieController {
	return &MovieController{}
}
