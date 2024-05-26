package dto

type CompanyDTO struct {
}

type CompanyCreateDTO struct {
	URL string `json:"url" validate:"required"`
}
