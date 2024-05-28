package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity/enum"
	"imdb_project/data/mapper"
	"net/http"
)

type ITvShowService interface {
	CreateTvShow(tvShow *dto.TvShowCreateDTO) dto.ResponseDTO[dto.TvShowDTO]
	LikeTvShow(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO]
	DislikeTvShow(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO]
	FindAllTvShow() dto.ResponseDTO[[]dto.TvShowDTO]
	FindTvShowById(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO]
	AddTvShowToWatchList(tvShowID uuid.UUID, watchListId uuid.UUID) dto.ResponseDTO[dto.WatchListDTO]
	RemoveTvShowFromWatchList(tvShowID uuid.UUID, watchListId uuid.UUID) dto.ResponseDTO[dto.WatchListDTO]
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

//...
