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
	Ratings     []entity.Rating    `json:"ratings"`
	Photos      []entity.Photo     `json:"photos"`
}

type MovieCreateDTO struct {
	ID        uuid.UUID          `json:"id"`
	Title     string             `json:"name"`
	Year      uint               `json:"year"`
	Trailers  []entity.Trailer   `json:"trailers"`
	Companies []entity.Company   `json:"companies"`
	Celebs    []entity.Celebrity `json:"celebs"`
	Ratings   []entity.Rating    `json:"ratings"`
	Photos    []entity.Photo     `json:"photos"`
}
