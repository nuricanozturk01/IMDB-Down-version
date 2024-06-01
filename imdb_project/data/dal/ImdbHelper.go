package dal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/mapper"
	"imdb_project/util"
	"log"
)

func (serviceHelper *ServiceHelper) CreateMovie(movie *dto.MovieCreateDTO) *entity.Movie {

	movieEntity := mapper.MovieCreateDtoToMovie(movie)

	result, err := serviceHelper.MovieRepository.Create(&movieEntity)

	if err != nil {
		log.Printf("Failed to create movie: %v", err)
		return nil
	}

	return result
}

func (serviceHelper *ServiceHelper) CreateTvShow(tvShow *dto.TvShowCreateDTO) *entity.TVShow {

	tvShowEntity := mapper.TvShowCreateDtoToTvShow(tvShow)

	result, err := serviceHelper.TvShowRepository.Create(&tvShowEntity)

	if err != nil {
		log.Printf("Failed to create tv show: %v", err)
		return nil
	}

	return result
}

func (serviceHelper *ServiceHelper) FindAllMovies() []entity.Movie {

	movies, err := serviceHelper.MovieRepository.FindAllEager([]string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

	if err != nil {
		log.Println("Failed to find movies:", err)
		return nil
	}

	return movies
}

func (serviceHelper *ServiceHelper) FindAllTvShows() []entity.TVShow {
	tvShows, err := serviceHelper.TvShowRepository.FindAll()

	if err != nil {
		log.Println("Failed to find tv shows:", err)
		return nil
	}

	return tvShows
}

func (serviceHelper *ServiceHelper) FindAllCelebrities() []entity.Celebrity {
	celebrities, err := serviceHelper.CelebrityRepository.FindAllEager([]string{"Movies", "TVShows", "Photos"})

	if err != nil {
		log.Println("Failed to find celebrities:", err)
		return nil
	}

	return celebrities
}

func (serviceHelper *ServiceHelper) FindMovieByID(id uuid.UUID) *entity.Movie {
	movie, err := serviceHelper.MovieRepository.FindByIdEager(id, []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

	if err != nil {
		log.Println("Failed to find movie:", err)
		return nil
	}

	return &movie
}

func (serviceHelper *ServiceHelper) FindTvShowByID(id uuid.UUID) *entity.TVShow {
	tvShow, err := serviceHelper.TvShowRepository.FindByID(id)

	if err != nil {
		log.Println("Failed to find tv show:", err)
		return nil
	}

	return &tvShow
}

func (serviceHelper *ServiceHelper) FindCelebrityByID(id uuid.UUID) *entity.Celebrity {
	celebrity, err := serviceHelper.CelebrityRepository.FindByID(id)

	if err != nil {
		log.Println("Failed to find celebrity:", err)
		return nil
	}

	return &celebrity
}

func (serviceHelper *ServiceHelper) Search(keyword string) dto.SearchDTO {
	movies, err := serviceHelper.searchMovieByKeyword(keyword)
	util.CheckError(err, "Failed to search movie by keyword:")

	tvShows, err := serviceHelper.searchTvShowByKeyword(keyword)
	util.CheckError(err, "Failed to search tv show by keyword:")

	celebs, err := serviceHelper.searchCelebrityByKeyword(keyword)
	util.CheckError(err, "Failed to search celebrity by keyword:")

	return dto.SearchDTO{Keyword: keyword, Movies: movies, TvShows: tvShows, Celebs: celebs}
}

func (serviceHelper *ServiceHelper) Like(mediaID, userID uuid.UUID, mediaType string) bool {

	like := entity.Like{MediaID: mediaID, MediaType: mediaType, UserID: userID}

	_, err := serviceHelper.LikeRepository.Create(&like)

	if err != nil {
		log.Println("Failed to like movie:", err)
		return false
	}

	return true
}

func (serviceHelper *ServiceHelper) Unlike(mediaID, userID uuid.UUID, mediaType string) bool {

	like := entity.Like{MediaID: mediaID, MediaType: mediaType, UserID: userID}

	_, err := serviceHelper.LikeRepository.DeleteById(like.ID)

	if err != nil {
		log.Println("Failed to unlike movie:", err)
		return false
	}
	return true
}

func (serviceHelper *ServiceHelper) AddWatchList(userID, mediaID uuid.UUID, mediaType string) bool {

	watchList, err := serviceHelper.UserRepository.FindByIdEager(userID, []string{"WatchList"})
	if err != nil {
		log.Println("Failed to find user watch list:", err)
		return false
	}

	item, err := serviceHelper.createWatchListItem(watchList.WatchList.ID, mediaID, mediaType)
	if err != nil {
		log.Println("Failed to create watch list item:", err)
		return false
	}

	_, err = serviceHelper.WatchListItemRepository.Create(&item)
	if err != nil {
		log.Println("Failed to add item to watch list:", err)
		return false
	}

	return true

}

func (serviceHelper *ServiceHelper) RemoveWatchList(userID, mediaID uuid.UUID, mediaType string) bool {

	watchList := serviceHelper.FindWatchListByUserID(userID)

	item, err := serviceHelper.findWatchListItem(watchList, mediaID, mediaType)
	if err != nil {
		log.Println("Failed to find watch list item:", err)
		return false
	}

	_, err = serviceHelper.WatchListItemRepository.DeleteById(item.ID)
	if err != nil {
		log.Println("Failed to remove item from watch list:", err)
		return false
	}

	return true
}

func (serviceHelper *ServiceHelper) CreateUser(user entity.User) *entity.User {
	result, err := serviceHelper.UserRepository.Create(&user)

	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil
	}

	return result
}

func (serviceHelper *ServiceHelper) FindUserByID(userID string) *entity.User {
	id, err := uuid.Parse(userID)

	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
		return nil
	}

	user, err := serviceHelper.UserRepository.FindByIdEager(id, []string{"WatchList", "Rates", "Likes"})
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return nil
	}

	return &user
}

func (serviceHelper *ServiceHelper) FindAllUsers() []entity.User {

	users, err := serviceHelper.UserRepository.FindAll()

	if err != nil {
		log.Println("Failed to find users:", err)
		return nil
	}

	return users
}

func (serviceHelper *ServiceHelper) FindUserByEmail(email string) *entity.User {

	user, err := serviceHelper.UserRepository.FindOneByFilter(findByEmailCallback(email))

	if err != nil {
		log.Println("Failed to find user:", err)
		return nil
	}

	return user
}

func (serviceHelper *ServiceHelper) FindWatchListByUserID(id uuid.UUID) *entity.WatchList {
	watchList, _ := serviceHelper.WatchListRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", id.String())
	}, []string{"Items"})

	return watchList
}

//...
