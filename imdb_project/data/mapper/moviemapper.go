package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func MovieCreateDtoToMovie(movieCreateDTO *dto.MovieCreateDTO) entity.Movie {
	return entity.Movie{
		ID:        movieCreateDTO.ID,
		Name:      movieCreateDTO.Title,
		Year:      movieCreateDTO.Year,
		Trailers:  movieCreateDTO.Trailers,
		Companies: movieCreateDTO.Companies,
		Celebs:    movieCreateDTO.Celebs,
		Ratings:   movieCreateDTO.Ratings,
		Photos:    movieCreateDTO.Photos,
	}
}

func MovieToMovieDTO(movie *entity.Movie) dto.MovieDTO {
	return dto.MovieDTO{
		ID:          movie.ID,
		Name:        movie.Name,
		AverageRate: movie.AverageRate,
		Year:        movie.Year,
		Popularity:  movie.Popularity,
		ClickCount:  movie.ClickCount,
		Trailers:    movie.Trailers,
		Companies:   movie.Companies,
		Celebs:      movie.Celebs,
		Ratings:     movie.Ratings,
		Photos:      movie.Photos,
	}
}
