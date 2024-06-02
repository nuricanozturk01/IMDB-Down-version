package service

import (
	"github.com/allegro/bigcache/v3"
	"imdb_project/data/dal"
	"imdb_project/data/entity"
)

type IInformationService interface {
	FindAllCountries() []string
}

type InformationService struct {
	ServiceHelper *dal.ImdbHelper
	Cache         *bigcache.BigCache
}

func NewInformationService(serviceHelper *dal.ImdbHelper, cache *bigcache.BigCache) *InformationService {
	return &InformationService{ServiceHelper: serviceHelper, Cache: cache}
}

func (service *InformationService) FindAllCountries() []entity.Country {
	return service.ServiceHelper.FindAllCountries()
}

func (service *InformationService) FindCitiesByCountry(country string) []entity.City {
	return service.ServiceHelper.FindCitiesByCountry(country)
}
