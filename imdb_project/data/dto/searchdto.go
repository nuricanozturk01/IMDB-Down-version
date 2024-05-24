package dto

type SearchDTO struct {
	Keyword string         `json:"keyword"`
	Movies  []MovieDTO     `json:"movies"`
	TvShows []TvShowDTO    `json:"tv_shows"`
	Celebs  []CelebrityDTO `json:"celebs"`
}
