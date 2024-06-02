package dal

import (
	"encoding/json"
	"github.com/allegro/bigcache/v3"
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
	tvShows, err := serviceHelper.TvShowRepository.FindAllEager([]string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

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
	tvShow, err := serviceHelper.TvShowRepository.FindByIdEager(id, []string{"Trailers", "Companies", "Celebs", "Photos", "Likes"})

	if err != nil {
		log.Println("Failed to find tv show:", err)
		return nil
	}

	return &tvShow
}

func (serviceHelper *ServiceHelper) FindCelebrityByID(id uuid.UUID) *entity.Celebrity {
	celebrity, err := serviceHelper.CelebrityRepository.FindByIdEager(id, []string{"Movies", "TVShows", "Photos"})
	var movies []entity.Movie
	var tvShows []entity.TVShow

	for _, celeb := range celebrity.Movies {
		movie := serviceHelper.FindMovieByID(celeb.ID)
		movies = append(movies, *movie)
	}

	for _, celeb := range celebrity.TVShows {
		tvShow := serviceHelper.FindTvShowByID(celeb.ID)
		tvShows = append(tvShows, *tvShow)
	}

	celebrity.Movies = movies
	celebrity.TVShows = tvShows
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

func (serviceHelper *ServiceHelper) FindAllCountries() []entity.Country {
	// Check if countries are in cache
	countriesBytes, err := serviceHelper.Cache.Get("countries")
	if err == nil {
		var countries []entity.Country
		err = json.Unmarshal(countriesBytes, &countries)
		if err != nil {
			log.Println("Failed to unmarshal countries from cache:", err)
			return nil
		}
		return countries
	}

	// If countries are not in cache, fetch from database
	countries, err := serviceHelper.CountryRepository.FindAll()
	if err != nil {
		log.Println("Failed to find countries:", err)
		return nil
	}

	// Cache the countries
	countriesBytes, err = json.Marshal(countries)
	if err != nil {
		log.Println("Failed to marshal countries for cache:", err)
		return countries
	}
	err = serviceHelper.Cache.Set("countries", countriesBytes)
	if err != nil {
		log.Println("Failed to set countries in cache:", err)
	}

	return countries
}

func (serviceHelper *ServiceHelper) FindCitiesByCountry(country string) []entity.City {
	// check if country is in cache
	citiesBytes, err := serviceHelper.Cache.Get(country)
	if err == nil {
		var cities []entity.City
		err = json.Unmarshal(citiesBytes, &cities)
		if err != nil {
			log.Println("Failed to unmarshal cities from cache:", err)
			return nil
		}
		return cities
	}

	// If country is not in cache, fetch from database
	countryObj, _ := serviceHelper.CountryRepository.FindOneByFilterEager(func(db *gorm.DB) *gorm.DB {
		return db.Where("country_name = ?", country)
	}, []string{"Cities"})

	go cacheCountryAndCities(countryObj, serviceHelper.Cache)

	return countryObj.Cities
}

func cacheCountryAndCities(obj *entity.Country, cache *bigcache.BigCache) {
	data, err := json.Marshal(obj.Cities)
	if err != nil {
		log.Println("Error while marshalling countries:", err)
		return
	}
	err = cache.Set(obj.CountryName, data)

	if err != nil {
		log.Println("Error while setting cache:", err)
	}
}

//...
