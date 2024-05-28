package dto

import "github.com/google/uuid"

type UserCreateDTO struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=45"`
	LastName  string `json:"last_name" validate:"required,min=3,max=45"`
	Username  string `json:"username" validate:"required,min=3,max=45"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=45"`
	Picture   string `json:"picture"`
	Locale    string `json:"locale"`
	GoogleID  string `json:"google_id"`
}

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type UserFullDTO struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Photos    []PhotoDTO `json:"photos"`
	Likes     []LikeDTO  `json:"likes"`
}

type GoogleUserDTO struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}
