package dal

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imdb_project/data/dto"
	"imdb_project/data/entity"
	"imdb_project/data/mapper"
	"imdb_project/data/repository"
	"log"
	"os"
)

type ImdbHelper struct {
	db                  *gorm.DB
	movieRepository     repository.IGenericRepository[entity.Movie, uuid.UUID]
	tvShowRepository    repository.IGenericRepository[entity.TVShow, uuid.UUID]
	celebrityRepository repository.IGenericRepository[entity.Celebrity, uuid.UUID]
}

func New(db *gorm.DB) *ImdbHelper {
	return &ImdbHelper{
		db:                  db,
		movieRepository:     repository.NewGenericRepository[entity.Movie, uuid.UUID](db),
		tvShowRepository:    repository.NewGenericRepository[entity.TVShow, uuid.UUID](db),
		celebrityRepository: repository.NewGenericRepository[entity.Celebrity, uuid.UUID](db),
	}
}

func InitDb() (*gorm.DB, error) {
	databaseConnection := os.Getenv("DB_DSN")

	var err error

	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&entity.Celebrity{},
		&entity.Company{},
		&entity.Movie{},
		&entity.Photo{},
		&entity.Rating{},
		&entity.Trailer{},
		&entity.TVShow{},
		&entity.User{},
		&entity.Watchlist{},
		&entity.WatchListItem{},
	)

	if err != nil {
		log.Println("Failed to migrate database:", err)
		return nil, err
	}

	log.Println("Database migrated successfully")

	return db, nil
}

func (ih *ImdbHelper) FindAllMovies() ([]entity.Movie, error) {
	return ih.movieRepository.FindAll()
}

func (ih *ImdbHelper) FindMovieByID(id uuid.UUID) (entity.Movie, error) {
	return ih.movieRepository.FindByID(id)
}

func (ih *ImdbHelper) CreateMovie(movie dto.MovieCreateDTO) (dto.MovieDTO, error) {
	movieEntity := mapper.MovieCreateDtoToMovie(&movie)

	result, err := ih.movieRepository.Create(&movieEntity)

	if err != nil {
		log.Panic("Failed to create movie:", err)
		return dto.MovieDTO{}, err
	}

	return mapper.MovieToMovieDTO(result), nil
}

func (ih *ImdbHelper) CreateTvShow(tvShow dto.TvShowCreateDTO) (dto.TvShowDTO, error) {
	tvShowEntity := mapper.TvShowCreateDtoToTvShow(&tvShow)

	result, err := ih.tvShowRepository.Create(&tvShowEntity)

	if err != nil {
		log.Panic("Failed to create tv show:", err)
		return dto.TvShowDTO{}, err
	}

	return mapper.TvShowToTvShowDTO(result), nil
}

func (ih *ImdbHelper) FindAllTvShows() ([]dto.TvShowDTO, error) {
	tvShows, err := ih.tvShowRepository.FindAll()

	if err != nil {
		log.Panic("Failed to find tv shows:", err)
		return nil, err
	}

	var tvShowDTOs []dto.TvShowDTO

	for _, tvShow := range tvShows {
		tvShowDTOs = append(tvShowDTOs, mapper.TvShowToTvShowDTO(&tvShow))
	}

	return tvShowDTOs, nil
}

func (ih *ImdbHelper) FindTvShowByID(id uuid.UUID) (dto.TvShowDTO, error) {
	tvShow, err := ih.tvShowRepository.FindByID(id)

	if err != nil {
		log.Panic("Failed to find tv show:", err)
		return dto.TvShowDTO{}, err
	}

	return mapper.TvShowToTvShowDTO(&tvShow), nil
}

func (ih *ImdbHelper) SearchMovieByKeyword(keyword string) ([]dto.MovieDTO, error) {

	filter, err := ih.movieRepository.FindByFilter(searchQueryCallback(keyword))

	if err != nil {
		return nil, err
	}

	var movieDTOs []dto.MovieDTO

	for _, movie := range filter {
		movieDTOs = append(movieDTOs, mapper.MovieToMovieDTO(&movie))
	}
	return movieDTOs, nil
}

func (ih *ImdbHelper) SearchTvShowByKeyword(keyword string) ([]dto.TvShowDTO, error) {
	filter, err := ih.tvShowRepository.FindByFilter(searchQueryCallback(keyword))

	if err != nil {
		return nil, err
	}

	var tvShowDTOs []dto.TvShowDTO

	for _, tvShow := range filter {
		tvShowDTOs = append(tvShowDTOs, mapper.TvShowToTvShowDTO(&tvShow))
	}
	return tvShowDTOs, nil

}

func (ih *ImdbHelper) SearchCelebrityByKeyword(keyword string) ([]dto.CelebrityDTO, error) {
	filter, err := ih.celebrityRepository.FindByFilter(searchQueryCallback(keyword))

	if err != nil {
		return nil, err
	}

	var celebrityDTOs []dto.CelebrityDTO

	for _, celebrity := range filter {
		celebrityDTOs = append(celebrityDTOs, mapper.CelebrityToCelebrityDTO(&celebrity))
	}
	return celebrityDTOs, nil
}

func searchQueryCallback(keyword string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+keyword+"%")
	}
}

func (ih *ImdbHelper) Search(keyword string) dto.SearchDTO {
	movies, err := ih.SearchMovieByKeyword(keyword)
	if err != nil {
		log.Panic("Failed to search movie by keyword:", err)
	}
	tvShows, err := ih.SearchTvShowByKeyword(keyword)
	if err != nil {
		log.Panic("Failed to search tv show by keyword:", err)
	}

	celebs, err := ih.SearchCelebrityByKeyword(keyword)
	if err != nil {
		log.Panic("Failed to search celebrity by keyword:", err)
	}

	return dto.SearchDTO{
		Keyword: keyword,
		Movies:  movies,
		TvShows: tvShows,
		Celebs:  celebs,
	}
}
