package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity/enum"
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
	return service.ServiceHelper.CreateMovie(movie)
}

func (service *MovieService) LikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.Like(movieId, userId, enum.MovieType)
}

func (service *MovieService) DislikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.Unlike(movieId, userId, enum.MovieType)
}

func (service *MovieService) FindAllMovies() dto.ResponseDTO[[]dto.MovieDTO] {
	return service.ServiceHelper.FindAllMovies()
}

func (service *MovieService) FindMovieById(movieId uuid.UUID) dto.ResponseDTO[dto.MovieDTO] {
	return service.ServiceHelper.FindMovieByID(movieId)
}

func (service *MovieService) AddMovieToWatchList(movieId uuid.UUID, mediaID uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.AddWatchList(movieId, mediaID, enum.MovieType)
}

func (service *MovieService) RemoveMovieFromWatchList(movieId uuid.UUID, watchListId uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.RemoveWatchList(movieId, watchListId, enum.MovieType)
}

//...
