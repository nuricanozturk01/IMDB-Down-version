package service

import (
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/mapper"
	"imdb_project/util"
)

type IAuthenticationService interface {
	Login(email string, password string) dto.ResponseDTO[dto.AuthResponseDTO]
	Register(registerDTO dto.UserCreateDTO) dto.ResponseDTO[dto.AuthResponseDTO]
}

type AuthenticationService struct {
	AuthHelper *dal.ImdbHelper
}

func NewAuthenticationService(authHelper *dal.ImdbHelper) *AuthenticationService {
	return &AuthenticationService{AuthHelper: authHelper}
}

// Login Normal login operation with email and password
func (service *AuthenticationService) Login(email string, password string) dto.ResponseDTO[dto.AuthResponseDTO] {

	user := service.AuthHelper.FindUserByEmail(email)

	if user == nil {
		return dto.ResponseDTO[dto.AuthResponseDTO]{StatusCode: 404, Data: &dto.AuthResponseDTO{Token: ""}}
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return dto.ResponseDTO[dto.AuthResponseDTO]{StatusCode: 401, Data: &dto.AuthResponseDTO{Token: ""}}
	}

	return dto.ResponseDTO[dto.AuthResponseDTO]{StatusCode: 200, Data: &dto.AuthResponseDTO{Token: "token"}}
}

func (service *AuthenticationService) LoginOAuth2(googleDTO *dto.GoogleUserDTO) dto.ResponseDTO[dto.UserDTO] {

	isExistUser := service.AuthHelper.FindUserByEmail(googleDTO.Email)

	if isExistUser == nil {
		return service.Register(mapper.GoogleUserToUserCreateDTO(googleDTO))
	}

	return dto.ResponseDTO[dto.UserDTO]{Message: "Success!", StatusCode: 200, Data: nil}
}

func (service *AuthenticationService) Register(registerDTO *dto.UserCreateDTO) dto.ResponseDTO[dto.UserDTO] {

	user := service.AuthHelper.FindUserByEmail(registerDTO.Email)

	if user != nil {
		return dto.ResponseDTO[dto.UserDTO]{StatusCode: 409, Data: nil}
	}

	createdUser := service.AuthHelper.CreateUser(mapper.UserCreateDTOToUser(registerDTO))

	if createdUser == nil {
		return dto.ResponseDTO[dto.UserDTO]{StatusCode: 500, Data: nil}
	}

	userDTO := mapper.UserToUserDTO(createdUser)

	return dto.ResponseDTO[dto.UserDTO]{StatusCode: 201, Data: &userDTO}
}
