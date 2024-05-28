package service

import (
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"net/http"
)

type ISearchService interface {
	Search(query string) dto.ResponseDTO[dto.SearchDTO]
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
