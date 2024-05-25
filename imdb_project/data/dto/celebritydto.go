package dto

import "github.com/google/uuid"

type CelebrityDTO struct {
	ID      uuid.UUID
	Name    string
	Movies  []MovieDTO
	Photos  []PhotoDTO
	TVShows []TvShowDTO
}
