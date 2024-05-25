package controller

import "imdb_project/service"

type MovieController struct {
	MovieService service.IMovieService
}

func NewMovieController(movieService service.IMovieService) *MovieController {
	return &MovieController{MovieService: movieService}
}
