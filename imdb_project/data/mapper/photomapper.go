package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func PhotoToPhotoDTO(photo *entity.Photo) dto.PhotoDTO {
	return dto.PhotoDTO{
		ID:  photo.ID,
		URL: photo.URL,
	}
}
