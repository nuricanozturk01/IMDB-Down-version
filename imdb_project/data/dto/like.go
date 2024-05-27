package dto

import "github.com/google/uuid"

type LikeDTO struct {
	ID        uuid.UUID `json:"movie_id"`
	UserID    uuid.UUID `json:"user_id"`
	MediaID   uuid.UUID `json:"media_id"`
	MediaType string    `json:"media_type"`
}
