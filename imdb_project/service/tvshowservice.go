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

type ITvShowService interface {
	CreateTvShow(tvShow *dto.TvShowCreateDTO) dto.ResponseDTO[dto.TvShowDTO]
	LikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool]
	DislikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool]
	FindAllTvShow() dto.ResponseDTO[[]dto.TvShowDTO]
	FindTvShowById(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO]
	AddTvShowToWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool]
	RemoveTvShowFromWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool]
	RateTvShow(tvShowId uuid.UUID, userId uuid.UUID, rate float64) dto.ResponseDTO[string]
}

type TvShowService struct {
	ServiceHelper *dal.ImdbHelper
}

func NewTvShowService(serviceHelper *dal.ImdbHelper) *TvShowService {
	return &TvShowService{ServiceHelper: serviceHelper}
}

func (service *TvShowService) CreateTvShow(tvShow *dto.TvShowCreateDTO) dto.ResponseDTO[dto.TvShowDTO] {
	tvShowEntity := service.ServiceHelper.CreateTvShow(tvShow)

	tvShowDTO := mapper.TvShowToTvShowDTO(tvShowEntity)

	return dto.ResponseDTO[dto.TvShowDTO]{Message: "Tv show created successfully", StatusCode: http.StatusCreated, Data: &tvShowDTO}
}

func (service *TvShowService) LikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool] {
	result := service.ServiceHelper.Like(tvShowID, userID, enum.TvShowType)
	return dto.ResponseDTO[bool]{Message: "Movie liked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *TvShowService) DislikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool] {
	result := service.ServiceHelper.Unlike(tvShowID, userID, enum.TvShowType)
	return dto.ResponseDTO[bool]{Message: "Movie unliked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *TvShowService) FindAllTvShow() dto.ResponseDTO[[]dto.TvShowDTO] {
	tvShows := service.ServiceHelper.FindAllTvShows()

	var tvShowDTOs []dto.TvShowDTO

	for _, tvShow := range tvShows {
		tvShowDTOs = append(tvShowDTOs, mapper.TvShowToTvShowDTO(&tvShow))
	}

	return dto.ResponseDTO[[]dto.TvShowDTO]{Message: "Tv shows fetched successfully", StatusCode: http.StatusOK, Data: &tvShowDTOs}
}

func (service *TvShowService) FindTvShowById(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO] {
	tvShow := service.ServiceHelper.FindTvShowByID(tvShowID)

	tvShowDTO := mapper.TvShowToTvShowDTO(tvShow)

	return dto.ResponseDTO[dto.TvShowDTO]{Message: "Tv show fetched successfully", StatusCode: http.StatusOK, Data: &tvShowDTO}
}

func (service *TvShowService) AddTvShowToWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool] {
	result := service.ServiceHelper.AddWatchList(tvShowID, watchListID, enum.TvShowType)

	return dto.ResponseDTO[bool]{Message: "Item added to watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *TvShowService) RemoveTvShowFromWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool] {

	result := service.ServiceHelper.RemoveWatchList(tvShowID, watchListID, enum.TvShowType)

	return dto.ResponseDTO[bool]{Message: "Item removed from watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *TvShowService) RateTvShow(tvShowId uuid.UUID, userId uuid.UUID, rate float64) dto.ResponseDTO[string] {

	var err error

	user, err := service.ServiceHelper.UserRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", userId.String())
	}, []string{"Rates", "Likes", "Photos"})

	for _, rt := range user.Rates {
		if rt.MediaID == tvShowId && rt.UserID == userId {
			return dto.ResponseDTO[string]{Message: "You have already rated this Tv Show", StatusCode: http.StatusConflict, Data: nil}
		}
	}

	// After checking if user has already rated the movie, we can proceed to rate the movie
	rate = math.Min(rate, 10)

	rateEntity := entity.Rate{MediaID: tvShowId, UserID: userId, Rate: rate, MediaType: enum.TvShowType}

	tvShow := service.ServiceHelper.FindTvShowByID(tvShowId)

	tvShow.AverageRate = math.Round((tvShow.AverageRate+rate)/2*100) / 100

	if _, err := service.ServiceHelper.TvShowRepository.Update(tvShow); err != nil {
		return dto.ResponseDTO[string]{Message: "Failed to rate Tv Show", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	if _, err = service.ServiceHelper.RateRepository.Create(&rateEntity); err != nil {
		return dto.ResponseDTO[string]{Message: "Failed to rate tv show", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	return dto.ResponseDTO[string]{Message: "Tv Show rated successfully", StatusCode: http.StatusOK, Data: nil}
}

//...
