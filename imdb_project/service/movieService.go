package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/entity/enum"
	"imdb_project/data/mapper"
	"math"
	"net/http"
)

type IMovieService interface {
	CreateMovie(movie *dto.MovieCreateDTO) dto.ResponseDTO[dto.MovieDTO]
	LikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool]
	DislikeMovie(movieId, userId uuid.UUID) dto.ResponseDTO[bool]
	FindAllMovies() dto.ResponseDTO[[]dto.MovieDTO]
	FindMovieById(movieId uuid.UUID) dto.ResponseDTO[dto.MovieDTO]
	AddMovieToWatchList(movieId uuid.UUID, mediaID uuid.UUID) dto.ResponseDTO[bool]
	RemoveMovieFromWatchList(movieId uuid.UUID, userId uuid.UUID) dto.ResponseDTO[bool]
	RateMovie(movieId uuid.UUID, userId uuid.UUID, rate float64) dto.ResponseDTO[string]
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

	movie, _ := service.ServiceHelper.MovieRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", movieId)
	}, []string{"Likes"})

	for _, like := range movie.Likes {
		if like.UserID == userId {
			return service.DislikeMovie(movieId, userId)
		}
	}

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

func (service *MovieService) AddMovieToWatchList(userId uuid.UUID, mediaID uuid.UUID) dto.ResponseDTO[bool] {

	watchList, _ := service.ServiceHelper.WatchListRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}, []string{"Items"})

	for _, item := range watchList.Items {
		if item.MediaID == mediaID {
			return dto.ResponseDTO[bool]{Message: "Media already exists in watch list!", StatusCode: http.StatusConflict, Data: nil}
		}
	}

	result := service.ServiceHelper.AddWatchList(userId, mediaID, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Item added to watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *MovieService) RemoveMovieFromWatchList(movieId uuid.UUID, userId uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.RemoveWatchList(userId, movieId, enum.MovieType)

	return dto.ResponseDTO[bool]{Message: "Item removed from watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *MovieService) RateMovie(movieId uuid.UUID, userId uuid.UUID, rate float64) dto.ResponseDTO[string] {

	var err error

	user, err := service.ServiceHelper.UserRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", userId.String())
	}, []string{"Rates", "Likes", "Photos"})

	for _, rt := range user.Rates {
		if rt.MediaID == movieId && rt.UserID == userId {
			return dto.ResponseDTO[string]{Message: "You have already rated this movie", StatusCode: http.StatusConflict, Data: nil}
		}
	}

	// After checking if user has already rated the movie, we can proceed to rate the movie
	rate = math.Min(rate, 10)

	rateEntity := entity.Rate{MediaID: movieId, UserID: userId, Rate: rate, MediaType: enum.MovieType}

	movie := service.ServiceHelper.FindMovieByID(movieId)
	if movie == nil {
		return dto.ResponseDTO[string]{Message: "Movie not found", StatusCode: http.StatusNotFound, Data: nil}
	}

	movie.AverageRate = math.Round((movie.AverageRate+rate)/2*100) / 100

	if _, err := service.ServiceHelper.MovieRepository.Update(movie); err != nil {
		return dto.ResponseDTO[string]{Message: "Failed to rate movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	if _, err = service.ServiceHelper.RateRepository.Create(&rateEntity); err != nil {
		return dto.ResponseDTO[string]{Message: "Failed to rate movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	return dto.ResponseDTO[string]{Message: "Movie rated successfully", StatusCode: http.StatusOK, Data: nil}
}

// ...
