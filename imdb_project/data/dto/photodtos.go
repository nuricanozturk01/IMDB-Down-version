package dto

type PhotoCreateDTO struct {
	URL string `json:"url" binding:"required"`
}

type PhotoDTO struct {
	ID        string `json:"id"`
	MediaID   string `json:"media_id"`
	MediaType string `json:"media_type"`
	URL       string `json:"url"`
}
