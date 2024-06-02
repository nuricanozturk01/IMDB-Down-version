package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/entity/enum"
	"net/http"
)

type ISearchService interface {
	Search(query string) dto.ResponseDTO[dto.SearchDTO]
	FindWatchList(id uuid.UUID) dto.ResponseDTO[dto.WatchListDTO]
}

type SearchService struct {
	ServiceHelper *dal.ImdbHelper
}

func NewSearchService(serviceHelper *dal.ImdbHelper) *SearchService {
	return &SearchService{ServiceHelper: serviceHelper}
}

func (service *SearchService) Search(query string) dto.ResponseDTO[dto.SearchDTO] {
	result := service.ServiceHelper.Search(query)

	return dto.ResponseDTO[dto.SearchDTO]{Message: "Search results fetched successfully", StatusCode: http.StatusOK, Data: &result}
}

func (service *SearchService) FindWatchList(id uuid.UUID) dto.ResponseDTO[dto.WatchListDTO] {
	watchList := service.ServiceHelper.FindWatchListByUserID(id)

	var movies []entity.Movie
	var tvShows []entity.TVShow

	for _, item := range watchList.Items {

		if item.MediaType == enum.MovieType {
			movie, _ := service.ServiceHelper.MovieRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
				return db.Where("id = ?", item.MediaID)
			}, []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

			movies = append(movies, *movie)
		} else if item.MediaType == enum.TvShowType {
			tvShow, _ := service.ServiceHelper.TvShowRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
				return db.Where("id = ?", item.MediaID)
			}, []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

			tvShows = append(tvShows, *tvShow)
		}
	}

	watchListDTO := &dto.WatchListDTO{ID: watchList.ID.String(), Movies: movies, TvShows: tvShows}

	return dto.ResponseDTO[dto.WatchListDTO]{Message: "Watch list fetched successfully", StatusCode: http.StatusOK, Data: watchListDTO}
}
