package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func UserToUserDTO(user *entity.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
func UserToUserDTOPtr(user *entity.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func UserCreateDTOToUser(userCreateDTO *dto.UserCreateDTO) entity.User {
	return entity.User{
		FirstName: userCreateDTO.FirstName,
		LastName:  userCreateDTO.LastName,
		Email:     userCreateDTO.Email,
		Password:  userCreateDTO.Password,
	}
}

func GoogleUserToUserCreateDTO(user *dto.GoogleUserDTO) *dto.UserCreateDTO {
	return &dto.UserCreateDTO{
		FirstName: user.GivenName,
		LastName:  user.FamilyName,
		Email:     user.Email,
		Picture:   user.Picture,
		Locale:    user.Locale,
		Password:  "from-google",
	}
}
