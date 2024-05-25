package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity/enum"
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
	return service.ServiceHelper.CreateTvShow(tvShow)
}

func (service *TvShowService) LikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.Like(tvShowID, userID, enum.TvShowType)
}

func (service *TvShowService) DislikeTvShow(tvShowID, userID uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.Unlike(tvShowID, userID, enum.TvShowType)
}

func (service *TvShowService) FindAllTvShow() dto.ResponseDTO[[]dto.TvShowDTO] {
	return service.ServiceHelper.FindAllTvShows()
}

func (service *TvShowService) FindTvShowById(tvShowID uuid.UUID) dto.ResponseDTO[dto.TvShowDTO] {
	return service.ServiceHelper.FindTvShowByID(tvShowID)
}

func (service *TvShowService) AddTvShowToWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.AddWatchList(tvShowID, watchListID, enum.TvShowType)
}

func (service *TvShowService) RemoveTvShowFromWatchList(tvShowID, watchListID uuid.UUID) dto.ResponseDTO[bool] {
	return service.ServiceHelper.RemoveWatchList(tvShowID, watchListID, enum.TvShowType)
}
