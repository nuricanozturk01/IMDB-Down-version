package dal

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/entity/enum"
	"imdb_project/data/mapper"
)

func searchQueryCallback(keyword string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+keyword+"%")
	}
}

func searchWatchListCallback(watchListID, mediaID uuid.UUID) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("watch_list_id = ? AND media_id = ?", watchListID, mediaID)
	}
}

func findByEmailCallback(email string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email = ?", email)
	}
}

func (serviceHelper *ServiceHelper) findWatchListItem(watchList *entity.WatchList, mediaID uuid.UUID, mediaType string) (*entity.WatchListItem, error) {
	switch mediaType {
	case enum.MovieType:
		for _, item := range watchList.Items {
			if item.MediaType == mediaType && item.MediaID == mediaID {
				return &item, nil
			}
		}

		return nil, fmt.Errorf("watch list item not found")

	case enum.TvShowType:
		for _, item := range watchList.Items {
			if item.MediaType == mediaType && item.MediaID == mediaID {
				return &item, nil
			}
		}

		return nil, fmt.Errorf("watch list item not found")
	default:
		return nil, fmt.Errorf("invalid media type: %s", mediaType)
	}
}

func (serviceHelper *ServiceHelper) createWatchListItem(watchListID, mediaID uuid.UUID, mediaType string) (entity.WatchListItem, error) {

	switch mediaType {
	case enum.MovieType:

		movie, _ := serviceHelper.MovieRepository.FindByID(mediaID)

		return entity.WatchListItem{WatchListID: watchListID, MediaID: movie.ID, MediaType: mediaType}, nil

	case enum.TvShowType:

		tvShow, _ := serviceHelper.TvShowRepository.FindByID(mediaID)

		return entity.WatchListItem{WatchListID: watchListID, MediaID: tvShow.ID, MediaType: mediaType}, nil

	default:
		return entity.WatchListItem{}, fmt.Errorf("invalid media type: %s", mediaType)
	}
}
func (serviceHelper *ServiceHelper) findWatchListItemByMediaID(watchListID, mediaID uuid.UUID) (entity.WatchListItem, error) {
	items, err := serviceHelper.WatchListItemRepository.FindByFilter(searchWatchListCallback(watchListID, mediaID))
	if err != nil {
		return entity.WatchListItem{}, fmt.Errorf("failed to find watch list item: %w", err)
	}
	if len(items) == 0 {
		return entity.WatchListItem{}, fmt.Errorf("watch list item not found")
	}
	return items[0], nil
}

func (serviceHelper *ServiceHelper) searchMovieByKeyword(keyword string) ([]dto.MovieDTO, error) {

	filter, err := serviceHelper.MovieRepository.FindByFilterEager(searchQueryCallback(keyword), []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

	if err != nil {
		return nil, err
	}

	var movieDTOs []dto.MovieDTO

	for _, movie := range filter {
		movieDTOs = append(movieDTOs, mapper.MovieToMovieDTO(&movie))
	}
	return movieDTOs, nil
}

func (serviceHelper *ServiceHelper) searchTvShowByKeyword(keyword string) ([]dto.TvShowDTO, error) {

	filter, err := serviceHelper.TvShowRepository.FindByFilterEager(searchQueryCallback(keyword), []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

	if err != nil {
		return nil, err
	}

	var tvShowDTOs []dto.TvShowDTO

	for _, tvShow := range filter {
		tvShowDTOs = append(tvShowDTOs, mapper.TvShowToTvShowDTO(&tvShow))
	}
	return tvShowDTOs, nil

}

func (serviceHelper *ServiceHelper) searchCelebrityByKeyword(keyword string) ([]dto.CelebrityDTO, error) {

	filter, err := serviceHelper.CelebrityRepository.FindByFilterEager(searchQueryCallback(keyword), []string{"Movies", "TVShows", "Photos"})

	if err != nil {
		return nil, err
	}

	var celebrityDTOs []dto.CelebrityDTO

	for _, celebrity := range filter {
		celebrityDTOs = append(celebrityDTOs, mapper.CelebrityToCelebrityDTO(&celebrity))
	}

	return celebrityDTOs, nil
}
