package tests

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"log"
	"os"
	"testing"
)

var appContext *AppContext
var userID uuid.UUID
var movieId uuid.UUID
var tvShowId uuid.UUID

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Successfully loaded .env file")

	// Reset the test database
	err = resetTestDB()
	if err != nil {
		log.Fatalf("Failed to reset database: %v", err)
	}

	appContext = configure()

	// Run the tests
	code := m.Run()

	// Clean up
	teardown()

	os.Exit(code)
}

func teardown() {
	err := appContext.DB.Migrator().DropTable(
		&entity.User{},
		&entity.Rate{},
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Like{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.WatchList{},
		&entity.WatchListItem{},
	)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	log.Println("Database tables dropped successfully")

}

func TestCreateUser(t *testing.T) {
	userCreateDTO := dto.UserCreateDTO{
		FirstName: "Nuri Can",
		LastName:  "Kurt",
		Email:     "canozturk309@gmail.com",
		Password:  "pass123",
		Picture:   "no_img.jpg",
		Locale:    "tr",
		GoogleID:  "123456789",
	}

	createdUser := appContext.UserService.CreateUser(&userCreateDTO)
	userID = createdUser.Data.ID

	if createdUser.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", createdUser.StatusCode)
	}

	if createdUser.Data.Email != userCreateDTO.Email {
		t.Errorf("Expected email %s, got %s", userCreateDTO.Email, createdUser.Data.Email)
	}

	if createdUser.Data.FirstName != userCreateDTO.FirstName {
		t.Errorf("Expected first name %s, got %s", userCreateDTO.FirstName, createdUser.Data.FirstName)
	}

	if createdUser.Data.LastName != userCreateDTO.LastName {
		t.Errorf("Expected last name %s, got %s", userCreateDTO.LastName, createdUser.Data.LastName)
	}
}

func TestLoginUser(t *testing.T) {
	loginDTO := dto.LoginDTO{
		Email:    "canozturk309@gmail.com",
		Password: "pass123",
	}

	loginResult := appContext.AuthenticationService.Login(loginDTO.Email, loginDTO.Password)

	if loginResult.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", loginResult.StatusCode)
	}

	if loginResult.Message != "Success!" {
		t.Errorf("Expected message 'Login successful', got %v", loginResult.Message)
	}

	if loginResult.Data.Email != loginDTO.Email {
		t.Errorf("Expected email %s, got %s", loginDTO.Email, loginResult.Data.Email)
	}

	if loginResult.Data.FirstName != "Nuri Can" {
		t.Errorf("Expected first name Nuri Can, got %s", loginResult.Data.FirstName)
	}
}

func TestCrateMovie(t *testing.T) {
	movieCreateDTO := dto.MovieCreateDTO{
		Name:        "The Shawshank Redemption",
		Description: "Two imprisoned",
		Year:        1994,
		Trailers: []entity.Trailer{
			{
				URL: "https://www.youtube.com/watch?v=6hB3S9bIaco",
			},
		},
		Photos: []entity.Photo{
			{
				URL: "https://www.imdb.com/title/tt0111161/mediaviewer/rm10105600",
			},
		},
		Companies: []entity.Company{
			{
				Name: "Castle Rock Entertainment",
			},
		},
		Celebs: []entity.Celebrity{
			{
				Name: "Tim Robbins",
			},
		},
	}

	createdMovie := appContext.MovieService.CreateMovie(&movieCreateDTO)
	movieId = createdMovie.Data.ID
	if createdMovie.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", createdMovie.StatusCode)
	}

	if createdMovie.Data.Name != movieCreateDTO.Name {
		t.Errorf("Expected name %s, got %s", movieCreateDTO.Name, createdMovie.Data.Name)
	}

	if createdMovie.Data.Description != movieCreateDTO.Description {
		t.Errorf("Expected description %s, got %s", movieCreateDTO.Description, createdMovie.Data.Description)
	}

	if createdMovie.Data.Year != movieCreateDTO.Year {
		t.Errorf("Expected year %d, got %d", movieCreateDTO.Year, createdMovie.Data.Year)
	}

	if createdMovie.Data.Trailers[0].URL != movieCreateDTO.Trailers[0].URL {
		t.Errorf("Expected trailer URL %s, got %s", movieCreateDTO.Trailers[0].URL, createdMovie.Data.Trailers[0].URL)
	}
}

func TestCreateTvShow(t *testing.T) {
	tvShowCreateDTO := dto.TvShowCreateDTO{
		Name:        "The Breaking Bad",
		Description: "A high school chemistry teacher turned meth",
		Year:        2008,
		Trailers: []entity.Trailer{
			{
				URL: "https://www.youtube.com/watch?v=HhesaQXLuRY",
			},
		},
		Photos: []entity.Photo{
			{
				URL: "https://www.imdb.com/title/tt0903747/mediaviewer/rm10105600",
			},
		},
		Companies: []entity.Company{
			{
				Name: "High Bridge Productions",
			},
		},
		Celebs: []entity.Celebrity{
			{
				Name: "Bryan Cranston",
			},
		},
	}

	createdTvShow := appContext.TvShowService.CreateTvShow(&tvShowCreateDTO)
	tvShowId = createdTvShow.Data.ID
	if createdTvShow.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", createdTvShow.StatusCode)
	}

	if createdTvShow.Data.Name != tvShowCreateDTO.Name {
		t.Errorf("Expected name %s, got %s", tvShowCreateDTO.Name, createdTvShow.Data.Name)
	}

	if createdTvShow.Data.Description != tvShowCreateDTO.Description {
		t.Errorf("Expected description %s, got %s", tvShowCreateDTO.Description, createdTvShow.Data.Description)
	}

	if createdTvShow.Data.Year != tvShowCreateDTO.Year {
		t.Errorf("Expected year %d, got %d", tvShowCreateDTO.Year, createdTvShow.Data.Year)
	}

	if createdTvShow.Data.Trailers[0].URL != tvShowCreateDTO.Trailers[0].URL {
		t.Errorf("Expected trailer URL %s, got %s", tvShowCreateDTO.Trailers[0].URL, createdTvShow.Data.Trailers[0].URL)
	}
}

func TestFindUserById(t *testing.T) {
	user := appContext.UserService.FindUserById(userID.String())

	if user.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", user.StatusCode)
	}

	if user.Data.ID != userID {
		t.Errorf("Expected user ID %s, got %s", userID, user.Data.ID)
	}
}

func TestRateMovie(t *testing.T) {
	ratedMovie := appContext.MovieService.RateMovie(movieId, userID, 5)

	if ratedMovie.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", ratedMovie.StatusCode)
	}

	if ratedMovie.Message != "Movie rated successfully" {
		t.Errorf("Expected message 'Movie rated successfully', got %v", ratedMovie.Data)
	}
}

func TestRateTvShow(t *testing.T) {
	ratedTvShow := appContext.TvShowService.RateTvShow(tvShowId, userID, 5)

	if ratedTvShow.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", ratedTvShow.StatusCode)
	}

	if ratedTvShow.Message != "Tv Show rated successfully" {
		t.Errorf("Expected message 'TV Show rated successfully', got %v", ratedTvShow.Data)
	}
}

func TestFindMovieById(t *testing.T) {
	movie := appContext.MovieService.FindMovieById(movieId)

	if movie.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", movie.StatusCode)
	}

	if movie.Data.ID != movieId {
		t.Errorf("Expected movie ID %s, got %s", movieId, movie.Data.ID)
	}
}

func TestFindTvShowById(t *testing.T) {
	tvShow := appContext.TvShowService.FindTvShowById(tvShowId)

	if tvShow.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", tvShow.StatusCode)
	}

	if tvShow.Data.ID != tvShowId {
		t.Errorf("Expected TV Show ID %s, got %s", tvShowId, tvShow.Data.ID)
	}
}

func TestMovieAddWatchList(t *testing.T) {
	result := appContext.MovieService.AddMovieToWatchList(userID, movieId)

	if result.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", result.StatusCode)
	}

	if result.Message != "Item added to watch list successfully" {
		t.Errorf("Expected message 'Item added to watch list successfully', got %v", result.Data)
	}

	if !*result.Data {
		t.Errorf("Expected data true, got %v", result.Data)
	}
}

func TestTvShowAddWatchList(t *testing.T) {
	result := appContext.TvShowService.AddTvShowToWatchList(userID, tvShowId)

	if result.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", result.StatusCode)
	}

	if result.Message != "Item added to watch list successfully" {
		t.Errorf("Expected message 'Item added to watch list successfully', got %v", result.Data)
	}

	if !*result.Data {
		t.Errorf("Expected data true, got %v", result.Data)
	}
}

func TestMovieRemoveWatchList(t *testing.T) {
	result := appContext.MovieService.RemoveMovieFromWatchList(movieId, userID)

	if result.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", result.StatusCode)
	}

	if result.Message != "Item removed from watch list successfully" {
		t.Errorf("Expected message 'Item removed from watch list successfully', got %v", result.Data)
	}

	if !*result.Data {
		t.Errorf("Expected data true, got %v", result.Data)
	}
}

func TestTvShowRemoveWatchList(t *testing.T) {
	result := appContext.TvShowService.RemoveTvShowFromWatchList(userID, tvShowId)

	if result.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", result.StatusCode)
	}

	if result.Message != "Item removed from watch list successfully" {
		t.Errorf("Expected message 'Item removed from watch list successfully', got %v", result.Data)
	}

	if !*result.Data {
		t.Errorf("Expected data true, got %v", result.Data)
	}
}

func TestSearchOperation(t *testing.T) {
	searchResult := appContext.SearchService.Search("The")
	// For the breaking bad and the shawshank redemption
	// The is added to breaking bad because of the test data

	if searchResult.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", searchResult.StatusCode)
	}

	if searchResult.Message != "Search results fetched successfully" {
		t.Errorf("Expected message 'Search results fetched successfully', got %v", searchResult.Message)
	}

	if len(searchResult.Data.Movies) == 0 {
		t.Errorf("Expected movies length greater than 0, got %d", len(searchResult.Data.Movies))
	}

	if len(searchResult.Data.TvShows) == 0 {
		t.Errorf("Expected tv shows length greater than 0, got %d", len(searchResult.Data.TvShows))
	}
}
