package service

import (
	"github.com/google/uuid"
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/mapper"
	"net/http"
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
	celebrity := service.ServiceHelper.FindCelebrityByID(id)

	celebrityDTO := mapper.CelebrityToCelebrityDTO(celebrity)

	return dto.ResponseDTO[dto.CelebrityDTO]{Message: "Celebrity fetched successfully", StatusCode: http.StatusOK, Data: &celebrityDTO}
}

func (service *CelebrityService) FindAllCelebrities() dto.ResponseDTO[[]dto.CelebrityDTO] {
	celebs := service.ServiceHelper.FindAllCelebrities()

	var celebrityDTOs []dto.CelebrityDTO

	for _, celebrity := range celebs {
		celebrityDTOs = append(celebrityDTOs, mapper.CelebrityToCelebrityDTO(&celebrity))
	}

	return dto.ResponseDTO[[]dto.CelebrityDTO]{Message: "Celebrities fetched successfully", StatusCode: http.StatusOK, Data: &celebrityDTOs}
}

//...
