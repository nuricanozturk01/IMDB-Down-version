package service

import (
	"imdb_project/data/dal"
	"imdb_project/data/dto"
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
	return service.ServiceHelper.Search(query)
}
