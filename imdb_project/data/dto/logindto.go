package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponseDTO struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
}
