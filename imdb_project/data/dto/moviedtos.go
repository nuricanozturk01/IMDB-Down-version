package dto

import (
	"github.com/google/uuid"
	"imdb_project/data/entity"
)

type MovieDTO struct {
	ID          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	AverageRate float64            `json:"average_rate"`
	Year        uint               `json:"year"`
	Popularity  uint               `json:"popularity"`
	ClickCount  uint               `json:"click_count"`
	Trailers    []entity.Trailer   `json:"trailers"`
	Companies   []entity.Company   `json:"companies"`
	Celebs      []entity.Celebrity `json:"celebs"`
	Likes       []entity.Like      `json:"likes"`
	Photos      []entity.Photo     `json:"photos"`
}

type MovieCreateDTO struct {
	Name      string             `json:"name" validate:"required"`
	Year      uint               `json:"year" validate:"required"`
	Trailers  []entity.Trailer   `json:"trailers" validate:"required"`
	Companies []entity.Company   `json:"companies" validate:"required"`
	Celebs    []entity.Celebrity `json:"celebs" validate:"required"`
	Likes     []entity.Like      `json:"likes"`
	Photos    []entity.Photo     `json:"photos" validate:"required"`
}
