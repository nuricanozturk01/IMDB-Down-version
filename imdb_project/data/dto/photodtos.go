package dto

import "github.com/google/uuid"

type PhotoCreateDTO struct {
	URL string `json:"url" validate:"required"`
}

type PhotoDTO struct {
	ID        uuid.UUID `json:"id"`
	MediaID   string    `json:"media_id"`
	MediaType string    `json:"media_type"`
	URL       string    `json:"url"`
}
