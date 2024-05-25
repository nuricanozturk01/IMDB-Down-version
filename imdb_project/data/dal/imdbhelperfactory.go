package dal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"imdb_project/data/entity"
	"imdb_project/data/repository"
)

type ServiceHelper = ImdbHelper

type ImdbHelper struct {
	Db                      *gorm.DB
	MovieRepository         repository.IGenericRepository[entity.Movie, uuid.UUID]
	TvShowRepository        repository.IGenericRepository[entity.TVShow, uuid.UUID]
	CelebrityRepository     repository.IGenericRepository[entity.Celebrity, uuid.UUID]
	WatchListRepository     repository.IGenericRepository[entity.WatchList, uuid.UUID]
	WatchListItemRepository repository.IGenericRepository[entity.WatchListItem, uuid.UUID]
	LikeRepository          repository.IGenericRepository[entity.Like, uuid.UUID]
	UserRepository          repository.IGenericRepository[entity.User, uuid.UUID]
}

func New(db *gorm.DB) *ImdbHelper {
	return &ImdbHelper{
		Db:                      db,
		MovieRepository:         repository.NewGenericRepository[entity.Movie, uuid.UUID](db),
		TvShowRepository:        repository.NewGenericRepository[entity.TVShow, uuid.UUID](db),
		CelebrityRepository:     repository.NewGenericRepository[entity.Celebrity, uuid.UUID](db),
		WatchListRepository:     repository.NewGenericRepository[entity.WatchList, uuid.UUID](db),
		WatchListItemRepository: repository.NewGenericRepository[entity.WatchListItem, uuid.UUID](db),
		LikeRepository:          repository.NewGenericRepository[entity.Like, uuid.UUID](db),
		UserRepository:          repository.NewGenericRepository[entity.User, uuid.UUID](db),
	}
}
