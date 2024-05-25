package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
)

type ICelebrityService interface {
	FindCelebrityByID(id uuid.UUID) dto.ResponseDTO[dto.CelebrityDTO]
	FindAllCelebrities() dto.ResponseDTO[[]dto.CelebrityDTO]
}

type CelebrityService struct {
	ServiceHelper *dal.ImdbHelper
}

func NewCelebrityService(serviceHelper *dal.ImdbHelper) *CelebrityService {
	return &CelebrityService{ServiceHelper: serviceHelper}
}

func (service *CelebrityService) FindCelebrityByID(id uuid.UUID) dto.ResponseDTO[dto.CelebrityDTO] {
	return service.ServiceHelper.FindCelebrityByID(id)
}

func (service *CelebrityService) FindAllCelebrities() dto.ResponseDTO[[]dto.CelebrityDTO] {
	return service.ServiceHelper.FindAllCelebrities()
}

//...
