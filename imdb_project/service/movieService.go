package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity/enum"
	"imdb_project/data/mapper"
	"net/http"
)

type IMovieService interface {
	CreateMovie(movie *dto.MovieCreateDTO) dto.ResponseDTO[dto.MovieDTO]
	LikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool]
	DislikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool]
	FindAllMovies() dto.ResponseDTO[[]dto.MovieDTO]
	FindMovieById(movieId uuid.UUID) dto.ResponseDTO[dto.MovieDTO]
	AddMovieToWatchList(movieId uuid.UUID, mediaID uuid.UUID) dto.ResponseDTO[bool]
	RemoveMovieFromWatchList(movieId uuid.UUID, watchListId uuid.UUID) dto.ResponseDTO[bool]
}

type MovieService struct {
	ServiceHelper *dal.ImdbHelper
}

func NewMovieService(serviceHelper *dal.ImdbHelper) *MovieService {
	return &MovieService{ServiceHelper: serviceHelper}
}

func (service *MovieService) CreateMovie(movie *dto.MovieCreateDTO) dto.ResponseDTO[dto.MovieDTO] {

	movieEntity := service.ServiceHelper.CreateMovie(movie)

	movieDTO := mapper.MovieToMovieDTO(movieEntity)

	return dto.ResponseDTO[dto.MovieDTO]{Message: "Movie created successfully", StatusCode: http.StatusCreated, Data: &movieDTO}
}

func (service *MovieService) LikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.Like(movieId, userId, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Movie liked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *MovieService) DislikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.Unlike(movieId, userId, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Movie unliked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *MovieService) FindAllMovies() dto.ResponseDTO[[]dto.MovieDTO] {

	movies := service.ServiceHelper.FindAllMovies()

	moviesDTO := mapper.ToMoviesDTO(movies)

	return dto.ResponseDTO[[]dto.MovieDTO]{Message: "Movies fetched successfully", StatusCode: http.StatusOK, Data: &moviesDTO}
}

func (service *MovieService) FindMovieById(movieId uuid.UUID) dto.ResponseDTO[dto.MovieDTO] {

	movie := service.ServiceHelper.FindMovieByID(movieId)

	movieDTO := mapper.MovieToMovieDTO(movie)

	return dto.ResponseDTO[dto.MovieDTO]{Message: "Movie fetched successfully", StatusCode: http.StatusOK, Data: &movieDTO}
}

func (service *MovieService) AddMovieToWatchList(movieId uuid.UUID, mediaID uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.AddWatchList(movieId, mediaID, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Item added to watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *MovieService) RemoveMovieFromWatchList(movieId uuid.UUID, watchListId uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.RemoveWatchList(movieId, watchListId, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Item removed from watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

//...
