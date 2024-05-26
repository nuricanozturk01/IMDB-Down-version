package dto

import "github.com/google/uuid"

type CelebrityDTO struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Movies  []MovieDTO  `json:"movies"`
	Photos  []PhotoDTO  `json:"photos"`
	TVShows []TvShowDTO `json:"tv_shows"`
}
