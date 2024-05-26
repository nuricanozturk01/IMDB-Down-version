package service

import (
	"imdb_project/data/dal"
	"imdb_project/data/dto"
)

type IUserService interface {
	CreateUser(user *dto.UserCreateDTO) dto.ResponseDTO[dto.UserDTO]
	FindUserById(userID string) dto.ResponseDTO[dto.UserDTO]
	FindAllUsers() dto.ResponseDTO[[]dto.UserDTO]
	FindUserByUsername(username string) dto.ResponseDTO[dto.UserDTO]
	FindUserByEmail(email string) dto.ResponseDTO[dto.UserDTO]
}

type UserService struct {
	ServiceHelper *dal.ImdbHelper
}

func NewUserService(serviceHelper *dal.ImdbHelper) *UserService {
	return &UserService{ServiceHelper: serviceHelper}
}

func (service *UserService) CreateUser(user *dto.UserCreateDTO) dto.ResponseDTO[dto.UserDTO] {
	return service.ServiceHelper.CreateUser(user)
}

func (service *UserService) FindUserById(userID string) dto.ResponseDTO[dto.UserDTO] {
	return service.ServiceHelper.FindUserByID(userID)
}

func (service *UserService) FindAllUsers() dto.ResponseDTO[[]dto.UserDTO] {
	return service.ServiceHelper.FindAllUsers()
}

func (service *UserService) FindUserByUsername(username string) dto.ResponseDTO[dto.UserDTO] {
	return service.ServiceHelper.FindUserByUsername(username)
}

func (service *UserService) FindUserByEmail(email string) dto.ResponseDTO[dto.UserDTO] {
	return service.ServiceHelper.FindUserByEmail(email)
}

// ...
