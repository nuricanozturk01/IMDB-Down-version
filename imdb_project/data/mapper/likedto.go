package mapper

import (
	"imdb_project/data/dto"
	"imdb_project/data/entity"
)

func LikedToLikedDTO(liked *entity.Like) dto.LikeDTO {
	return dto.LikeDTO{
		ID:        liked.ID,
		UserID:    liked.UserID,
		MediaID:   liked.MediaID,
		MediaType: liked.MediaType,
	}
}
