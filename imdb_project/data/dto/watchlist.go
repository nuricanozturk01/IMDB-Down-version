package dto

import (
	"github.com/google/uuid"
	"imdb_project/data/entity"
)

type WatchListItemDTO[T entity.Movie | entity.TVShow] struct {
	ID     string `json:"id"`
	Medias []T    `json:"medias"`
}

type WatchListDTO struct {
	ID      string                            `json:"id"`
	Movies  []WatchListItemDTO[entity.Movie]  `json:"movies"`
	TvShows []WatchListItemDTO[entity.TVShow] `json:"tv_shows"`
}

type WatchListItemIoDTO struct {
	UserID  uuid.UUID `json:"user_id"`
	MediaID uuid.UUID `json:"media_id"`
}
