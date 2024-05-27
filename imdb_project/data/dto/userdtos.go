package dto

import "github.com/google/uuid"

type UserCreateDTO struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=45"`
	LastName  string `json:"last_name" validate:"required,min=3,max=45"`
	Username  string `json:"username" validate:"required,min=3,max=45"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=45"`
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
