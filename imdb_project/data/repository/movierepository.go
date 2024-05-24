package repository

import (
	"gorm.io/gorm"
	"imdb_project/data/entity"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (mr *MovieRepository) GetMovieByID(id string) (*entity.Movie, error) {
	var movie entity.Movie
	err := mr.db.Where("movie_id = ?", id).First(&movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (mr *MovieRepository) GetMovies() ([]entity.Movie, error) {
	var movies []entity.Movie
	err := mr.db.Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (mr *MovieRepository) CreateMovie(movie entity.Movie) error {
	err := mr.db.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *MovieRepository) UpdateMovie(id string, movie entity.Movie) error {
	err := mr.db.Where("movie_id = ?", id).Updates(&movie).Error
	if err != nil {
		return err
	}
	return nil
}
