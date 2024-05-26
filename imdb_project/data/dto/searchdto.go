package dto

type SearchDTO struct {
	Keyword string         `json:"keyword" validate:"required"`
	Movies  []MovieDTO     `json:"movies" validate:"required"`
	TvShows []TvShowDTO    `json:"tv_shows" validate:"required"`
	Celebs  []CelebrityDTO `json:"celebs" validate:"required"`
}
