package dto

type ResponseDTO[T any] struct {
	Message    string `json:"message"`
	StatusCode uint32 `json:"status_code"`
	Data       *T     `json:"data"`
}
