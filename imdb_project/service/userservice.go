package service

import (
	"imdb_project/data/dal"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/mapper"
	"imdb_project/util"
	"log"
	"net/http"
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
	// to user entity
	userEntity := mapper.UserCreateDTOToUser(user)
	// hash password
	hashedPassword, err := util.HashPassword(user.Password)

	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return dto.ResponseDTO[dto.UserDTO]{Message: "Failed to hash password", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	// set hashed password
	userEntity.Password = hashedPassword

	// create user
	createdUser := service.ServiceHelper.CreateUser(userEntity)

	// to user dto
	userDTO := mapper.UserToUserDTO(createdUser)

	return dto.ResponseDTO[dto.UserDTO]{Message: "User created successfully", StatusCode: http.StatusCreated, Data: &userDTO}
}

func (service *UserService) FindUserById(userID string) dto.ResponseDTO[dto.UserDTO] {
	user := service.ServiceHelper.FindUserByID(userID)

	if user == nil {
		return dto.ResponseDTO[dto.UserDTO]{Message: "User not found", StatusCode: http.StatusNotFound, Data: nil}
	}

	userDTO := mapper.UserToUserDTO(user)

	return dto.ResponseDTO[dto.UserDTO]{Message: "User fetched successfully", StatusCode: http.StatusOK, Data: &userDTO}
}

func (service *UserService) FindAllUsers() dto.ResponseDTO[[]dto.UserDTO] {
	users := service.ServiceHelper.FindAllUsers()

	var userDTOs []dto.UserDTO

	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToUserDTO(&user))
	}

	return dto.ResponseDTO[[]dto.UserDTO]{Message: "Users fetched successfully", StatusCode: http.StatusOK, Data: &userDTOs}
}

func (service *UserService) FindUserByUsername(username string) dto.ResponseDTO[dto.UserDTO] {
	user := service.ServiceHelper.FindUserByUsername(username)

	userDTO := mapper.UserToUserDTO(user)

	return dto.ResponseDTO[dto.UserDTO]{Message: "User fetched successfully", StatusCode: http.StatusOK, Data: &userDTO}
}

func (service *UserService) FindUserByEmail(email string) dto.ResponseDTO[dto.UserDTO] {
	user := service.ServiceHelper.FindUserByEmail(email)

	userDTO := mapper.UserToUserDTO(user)

	return dto.ResponseDTO[dto.UserDTO]{Message: "User fetched successfully", StatusCode: http.StatusOK, Data: &userDTO}
}

func (service *UserService) FindUserByEmailForAuth(email string) *entity.User {
	return service.ServiceHelper.FindUserByEmail(email)
}

// ...
