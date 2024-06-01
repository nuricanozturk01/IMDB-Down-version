package dto

import (
	"github.com/google/uuid"
	"imdb_project/data/entity"
)

type CelebrityDTO struct {
	ID      uuid.UUID       `json:"id"`
	Name    string          `json:"name"`
	Movies  []entity.Movie  `json:"movies"`
	Photos  []entity.Photo  `json:"photos"`
	TVShows []entity.TVShow `json:"tv_shows"`
}
