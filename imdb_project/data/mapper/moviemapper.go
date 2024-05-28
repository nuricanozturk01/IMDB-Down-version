package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func MovieCreateDtoToMovie(movieCreateDTO *dto.MovieCreateDTO) entity.Movie {
	return entity.Movie{
		Name:        movieCreateDTO.Name,
		AverageRate: 0,
		Year:        movieCreateDTO.Year,
		Popularity:  0,
		ClickCount:  0,
		Trailers:    movieCreateDTO.Trailers,
		Companies:   movieCreateDTO.Companies,
		Celebs:      movieCreateDTO.Celebs,
		Likes:       movieCreateDTO.Likes,
		Photos:      movieCreateDTO.Photos,
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
		Likes:       movie.Likes,
		Photos:      movie.Photos,
	}
}

func ToMoviesDTO(movies []entity.Movie) []dto.MovieDTO {
	var movieDTOs []dto.MovieDTO

	for _, movie := range movies {
		movieDTOs = append(movieDTOs, MovieToMovieDTO(&movie))
	}

	return movieDTOs
}
