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

func UserCreateDTOToUser(userCreateDTO *dto.UserCreateDTO) entity.User {
	return entity.User{
		FirstName: userCreateDTO.FirstName,
		LastName:  userCreateDTO.LastName,
		Email:     userCreateDTO.Email,
		Password:  userCreateDTO.Password,
	}
}

func UserToUserFullDTO(user *entity.User) dto.UserFullDTO {
	photos := make([]dto.PhotoDTO, 0)
	likes := make([]dto.LikeDTO, 0)

	for _, photo := range user.Photos {
		photos = append(photos, PhotoToPhotoDTO(&photo))
	}

	for _, like := range user.Likes {
		likes = append(likes, LikedToLikedDTO(&like))
	}

	return dto.UserFullDTO{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Photos:    photos,
		Likes:     likes,
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
