package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func TvShowCreateDtoToTvShow(tvShowCreateDTO *dto.TvShowCreateDTO) entity.TVShow {
	return entity.TVShow{
		Name:         tvShowCreateDTO.Name,
		Year:         tvShowCreateDTO.Year,
		EpisodeCount: tvShowCreateDTO.EpisodeCount,
		SeasonCount:  tvShowCreateDTO.SeasonCount,
		Photos:       tvShowCreateDTO.Photos,
		Trailers:     tvShowCreateDTO.Trailers,
		Companies:    tvShowCreateDTO.Companies,
		Celebs:       tvShowCreateDTO.Celebs,
	}
}

func TvShowToTvShowDTO(tvShow *entity.TVShow) dto.TvShowDTO {
	return dto.TvShowDTO{
		ID:           tvShow.ID,
		Name:         tvShow.Name,
		Year:         tvShow.Year,
		Popularity:   tvShow.Popularity,
		AverageRate:  tvShow.AverageRate,
		ClickCount:   tvShow.ClickCount,
		EpisodeCount: tvShow.EpisodeCount,
		SeasonCount:  tvShow.SeasonCount,
		Photos:       tvShow.Photos,
		Trailers:     tvShow.Trailers,
		Companies:    tvShow.Companies,
		Celebs:       tvShow.Celebs,
		//Ratings:      tvShow.Ratings,
	}
}
