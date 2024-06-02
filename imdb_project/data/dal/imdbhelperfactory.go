package dal

import (
	"github.com/allegro/bigcache/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/entity"
	"imdb_project/data/repository"
)

type ServiceHelper = ImdbHelper

type ImdbHelper struct {
	Db                      *gorm.DB
	Cache                   *bigcache.BigCache
	MovieRepository         repository.IGenericRepository[entity.Movie, uuid.UUID]
	TvShowRepository        repository.IGenericRepository[entity.TVShow, uuid.UUID]
	LikeRepository          repository.IGenericRepository[entity.Like, uuid.UUID]
	UserRepository          repository.IGenericRepository[entity.User, uuid.UUID]
	RateRepository          repository.IGenericRepository[entity.Rate, uuid.UUID]
	CelebrityRepository     repository.IGenericRepository[entity.Celebrity, uuid.UUID]
	WatchListRepository     repository.IGenericRepository[entity.WatchList, uuid.UUID]
	WatchListItemRepository repository.IGenericRepository[entity.WatchListItem, uuid.UUID]
	CountryRepository       repository.IGenericRepository[entity.Country, string]
	CityRepository          repository.IGenericRepository[entity.City, int]
}

func New(db *gorm.DB, cache *bigcache.BigCache) *ImdbHelper {
	return &ImdbHelper{
		Db:                      db,
		Cache:                   cache,
		MovieRepository:         repository.NewGenericRepository[entity.Movie, uuid.UUID](db),
		TvShowRepository:        repository.NewGenericRepository[entity.TVShow, uuid.UUID](db),
		LikeRepository:          repository.NewGenericRepository[entity.Like, uuid.UUID](db),
		UserRepository:          repository.NewGenericRepository[entity.User, uuid.UUID](db),
		RateRepository:          repository.NewGenericRepository[entity.Rate, uuid.UUID](db),
		CelebrityRepository:     repository.NewGenericRepository[entity.Celebrity, uuid.UUID](db),
		WatchListRepository:     repository.NewGenericRepository[entity.WatchList, uuid.UUID](db),
		WatchListItemRepository: repository.NewGenericRepository[entity.WatchListItem, uuid.UUID](db),
		CountryRepository:       repository.NewGenericRepository[entity.Country, string](db),
		CityRepository:          repository.NewGenericRepository[entity.City, int](db),
	}
}
