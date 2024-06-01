package dto

import (
	"imdb_project/data/entity"
)

type WatchListDTO struct {
	ID      string          `json:"id"`
	Movies  []entity.Movie  `json:"movies"`
	TvShows []entity.TVShow `json:"tv_shows"`
}
