package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func CelebrityToCelebrityDTO(celebrity *entity.Celebrity) dto.CelebrityDTO {
	return dto.CelebrityDTO{
		ID:   celebrity.ID,
		Name: celebrity.Name,
	}
}
