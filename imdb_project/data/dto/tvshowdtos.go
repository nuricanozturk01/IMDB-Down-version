package dto

import (
	"github.com/google/uuid"
	"imdb_project/data/entity"
)

type TvShowCreateDTO struct {
	Name         string             `json:"name" validate:"required"`
	Year         int                `json:"year" validate:"required"`
	EpisodeCount uint32             `json:"episode_count" validate:"required"`
	SeasonCount  uint32             `json:"season_count" validate:"required"`
	Photos       []entity.Photo     `json:"photos" validate:"required"`
	Trailers     []entity.Trailer   `json:"trailers" validate:"required"`
	Companies    []entity.Company   `json:"companies" validate:"required"`
	Celebs       []entity.Celebrity `json:"celebs" validate:"required"`
}

type TvShowDTO struct {
	ID           uuid.UUID          `json:"id"`
	Name         string             `json:"name"`
	Year         int                `json:"year"`
	Popularity   int                `json:"popularity"`
	AverageRate  float64            `json:"average_rate"`
	ClickCount   uint32             `json:"click_count"`
	EpisodeCount uint32             `json:"episode_count"`
	SeasonCount  uint32             `json:"season_count"`
	Photos       []entity.Photo     `json:"photos"`
	Trailers     []entity.Trailer   `json:"trailers"`
	Companies    []entity.Company   `json:"companies"`
	Celebs       []entity.Celebrity `json:"celebs"`
}
