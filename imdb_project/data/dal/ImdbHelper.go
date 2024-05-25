package dal

import (
	"github.com/google/uuid"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/mapper"
	"imdb_project/util"
	"log"
	"net/http"
)

func (serviceHelper *ServiceHelper) CreateMovie(movie *dto.MovieCreateDTO) dto.ResponseDTO[dto.MovieDTO] {

	movieEntity := mapper.MovieCreateDtoToMovie(movie)

	result, err := serviceHelper.MovieRepository.Create(&movieEntity)

	if err != nil {
		log.Printf("Failed to create movie: %v", err)
		return dto.ResponseDTO[dto.MovieDTO]{Message: "Failed to create movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	movieDTO := mapper.MovieToMovieDTO(result)

	return dto.ResponseDTO[dto.MovieDTO]{Message: "Movie created successfully", StatusCode: http.StatusCreated, Data: &movieDTO}
}

func (serviceHelper *ServiceHelper) CreateTvShow(tvShow *dto.TvShowCreateDTO) dto.ResponseDTO[dto.TvShowDTO] {

	tvShowEntity := mapper.TvShowCreateDtoToTvShow(tvShow)

	result, err := serviceHelper.TvShowRepository.Create(&tvShowEntity)

	if err != nil {
		log.Printf("Failed to create tv show: %v", err)
		return dto.ResponseDTO[dto.TvShowDTO]{Message: "Failed to create tv show", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	tvShowDTO := mapper.TvShowToTvShowDTO(result)

	return dto.ResponseDTO[dto.TvShowDTO]{Message: "Tv show created successfully", StatusCode: http.StatusCreated, Data: &tvShowDTO}
}

func (serviceHelper *ServiceHelper) FindAllMovies() dto.ResponseDTO[[]dto.MovieDTO] {
	movies, err := serviceHelper.MovieRepository.FindAll()

	if err != nil {
		log.Println("Failed to find movies:", err)
		return dto.ResponseDTO[[]dto.MovieDTO]{Message: "Failed to fetch movies", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	var movieDTOs []dto.MovieDTO

	for _, movie := range movies {
		movieDTOs = append(movieDTOs, mapper.MovieToMovieDTO(&movie))
	}

	return dto.ResponseDTO[[]dto.MovieDTO]{Message: "Movies fetched successfully", StatusCode: http.StatusOK, Data: &movieDTOs}
}

func (serviceHelper *ServiceHelper) FindAllTvShows() dto.ResponseDTO[[]dto.TvShowDTO] {
	tvShows, err := serviceHelper.TvShowRepository.FindAll()

	if err != nil {
		log.Println("Failed to find tv shows:", err)
		return dto.ResponseDTO[[]dto.TvShowDTO]{Message: "Failed to fetch tv shows", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	var tvShowDTOs []dto.TvShowDTO

	for _, tvShow := range tvShows {
		tvShowDTOs = append(tvShowDTOs, mapper.TvShowToTvShowDTO(&tvShow))
	}

	return dto.ResponseDTO[[]dto.TvShowDTO]{Message: "Tv shows fetched successfully", StatusCode: http.StatusOK, Data: &tvShowDTOs}
}

func (serviceHelper *ServiceHelper) FindAllCelebrities() dto.ResponseDTO[[]dto.CelebrityDTO] {
	celebrities, err := serviceHelper.CelebrityRepository.FindAll()

	if err != nil {
		log.Println("Failed to find celebrities:", err)
		return dto.ResponseDTO[[]dto.CelebrityDTO]{Message: "Failed to fetch celebrities", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	var celebrityDTOs []dto.CelebrityDTO

	for _, celebrity := range celebrities {
		celebrityDTOs = append(celebrityDTOs, mapper.CelebrityToCelebrityDTO(&celebrity))
	}

	return dto.ResponseDTO[[]dto.CelebrityDTO]{Message: "Celebrities fetched successfully", StatusCode: http.StatusOK, Data: &celebrityDTOs}
}

func (serviceHelper *ServiceHelper) FindMovieByID(id uuid.UUID) dto.ResponseDTO[dto.MovieDTO] {
	movie, err := serviceHelper.MovieRepository.FindByID(id)

	if err != nil {
		log.Println("Failed to find movie:", err)
		return dto.ResponseDTO[dto.MovieDTO]{Message: "Failed to fetch movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	movieDTO := mapper.MovieToMovieDTO(&movie)

	return dto.ResponseDTO[dto.MovieDTO]{Message: "Movie fetched successfully", StatusCode: http.StatusOK, Data: &movieDTO}
}

func (serviceHelper *ServiceHelper) FindTvShowByID(id uuid.UUID) dto.ResponseDTO[dto.TvShowDTO] {
	tvShow, err := serviceHelper.TvShowRepository.FindByID(id)

	if err != nil {
		log.Println("Failed to find tv show:", err)
		return dto.ResponseDTO[dto.TvShowDTO]{Message: "Failed to fetch tv show", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	tvShowDTO := mapper.TvShowToTvShowDTO(&tvShow)

	return dto.ResponseDTO[dto.TvShowDTO]{Message: "Tv show fetched successfully", StatusCode: http.StatusOK, Data: &tvShowDTO}
}

func (serviceHelper *ServiceHelper) FindCelebrityByID(id uuid.UUID) dto.ResponseDTO[dto.CelebrityDTO] {
	celebrity, err := serviceHelper.CelebrityRepository.FindByID(id)

	if err != nil {
		log.Println("Failed to find celebrity:", err)
		return dto.ResponseDTO[dto.CelebrityDTO]{Message: "Failed to fetch celebrity", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	celebrityDTO := mapper.CelebrityToCelebrityDTO(&celebrity)

	return dto.ResponseDTO[dto.CelebrityDTO]{Message: "Celebrity fetched successfully", StatusCode: http.StatusOK, Data: &celebrityDTO}
}

func (serviceHelper *ServiceHelper) Search(keyword string) dto.ResponseDTO[dto.SearchDTO] {
	movies, err := serviceHelper.searchMovieByKeyword(keyword)
	util.CheckError(err, "Failed to search movie by keyword:")

	tvShows, err := serviceHelper.searchTvShowByKeyword(keyword)
	util.CheckError(err, "Failed to search tv show by keyword:")

	celebs, err := serviceHelper.searchCelebrityByKeyword(keyword)
	util.CheckError(err, "Failed to search celebrity by keyword:")

	searchDTO := dto.SearchDTO{Movies: movies, TvShows: tvShows, Celebs: celebs}

	return dto.ResponseDTO[dto.SearchDTO]{Message: "Search results fetched successfully", StatusCode: http.StatusOK, Data: &searchDTO}
}

func (serviceHelper *ServiceHelper) Like(mediaID, userID uuid.UUID, mediaType string) dto.ResponseDTO[bool] {

	like := entity.Like{MediaID: mediaID, MediaType: mediaType, UserID: userID}

	_, err := serviceHelper.LikeRepository.Create(&like)

	if err != nil {
		log.Println("Failed to like movie:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to like movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	result := true
	return dto.ResponseDTO[bool]{Message: "Movie liked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (serviceHelper *ServiceHelper) Unlike(mediaID, userID uuid.UUID, mediaType string) dto.ResponseDTO[bool] {

	like := entity.Like{MediaID: mediaID, MediaType: mediaType, UserID: userID}

	_, err := serviceHelper.LikeRepository.DeleteById(like.ID)

	if err != nil {
		log.Println("Failed to unlike movie:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to unlike movie", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	result := false
	return dto.ResponseDTO[bool]{Message: "Movie unliked successfully", StatusCode: http.StatusOK, Data: &result}
}

func (serviceHelper *ServiceHelper) AddWatchList(userID, mediaID uuid.UUID, mediaType string) dto.ResponseDTO[bool] {

	watchList, err := serviceHelper.UserRepository.FindByIdEager(userID, []string{"WatchList"})
	if err != nil {
		log.Println("Failed to find user watch list:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to find user watch list", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	item, err := serviceHelper.createWatchListItem(watchList.WatchList.ID, mediaID, mediaType)
	if err != nil {
		log.Println("Failed to create watch list item:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to create watch list item", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	_, err = serviceHelper.WatchListItemRepository.Create(&item)
	if err != nil {
		log.Println("Failed to add item to watch list:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to add item to watch list", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	result := true
	return dto.ResponseDTO[bool]{Message: "Item added to watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

func (serviceHelper *ServiceHelper) RemoveWatchList(userID, mediaID uuid.UUID, mediaType string) dto.ResponseDTO[bool] {

	watchList, err := serviceHelper.UserRepository.FindByIdEager(userID, []string{"WatchList"})
	if err != nil {
		log.Println("Failed to find user watch list:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to find user watch list", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	item, err := serviceHelper.findWatchListItem(watchList.WatchList.ID, mediaID, mediaType)
	if err != nil {
		log.Println("Failed to find watch list item:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to find watch list item", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	_, err = serviceHelper.WatchListItemRepository.DeleteById(item.ID)
	if err != nil {
		log.Println("Failed to remove item from watch list:", err)
		return dto.ResponseDTO[bool]{Message: "Failed to remove item from watch list", StatusCode: http.StatusInternalServerError, Data: nil}
	}

	result := true
	return dto.ResponseDTO[bool]{Message: "Item removed from watch list successfully", StatusCode: http.StatusOK, Data: &result}
}

//...
