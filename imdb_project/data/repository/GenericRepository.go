package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type IGenericRepository[T any, R any] interface {
	FindAll() ([]T, error)
	FindAllEager(eagerAssociations []string) ([]T, error)
	FindByID(id R) (T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	DeleteById(id R) (bool, error)
	FindByFilter(predicate func(*gorm.DB) *gorm.DB) ([]T, error)
	FindOneByFilter(predicate func(*gorm.DB) *gorm.DB) (*T, error)
	FindByFilterEager(predicate func(*gorm.DB) *gorm.DB, eagerAssociations []string) ([]T, error)
	FindOneByFilterEager(predicate func(*gorm.DB) *gorm.DB, eagerAssociations []string) (*T, error)
	FindByIdEager(id R, eagerAssociations []string) (T, error)
}

type GenericRepository[T, R any] struct {
	Db *gorm.DB
}

func NewGenericRepository[T, R any](db *gorm.DB) IGenericRepository[T, R] {
	return &GenericRepository[T, R]{Db: db}
}
func (repository *GenericRepository[T, R]) FindAll() ([]T, error) {
	var entities []T
	if err := repository.Db.Find(&entities).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}
	return entities, nil
}

func (repository *GenericRepository[T, R]) FindByID(id R) (T, error) {
	var entity T
	if err := repository.Db.First(&entity, id).Error; err != nil {
		log.Panic("Error while fetching entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (repository *GenericRepository[T, R]) Create(entity *T) (*T, error) {
	if err := repository.Db.Create(&entity).Error; err != nil {
		//log.Panic("Error while creating entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (repository *GenericRepository[T, R]) Update(entity *T) (*T, error) {
	if err := repository.Db.Save(entity).Error; err != nil {
		log.Panic("Error while updating entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (repository *GenericRepository[T, R]) DeleteById(id R) (bool, error) {
	if err := repository.Db.Delete(new(T), id).Error; err != nil {
		log.Panic("Error while deleting entity: ", err)
		return false, err
	}
	return true, nil
}

func (repository *GenericRepository[T, R]) FindByFilter(predicate func(*gorm.DB) *gorm.DB) ([]T, error) {
	var entities []T
	query := predicate(repository.Db)
	if err := query.Find(&entities).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}
	return entities, nil
}

func (repository *GenericRepository[T, R]) FindOneByFilter(predicate func(*gorm.DB) *gorm.DB) (*T, error) {
	var entity T
	query := predicate(repository.Db)

	if err := query.First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("Error while fetching entity: %v", err)
		return nil, err
	}

	return &entity, nil
}

func (repository *GenericRepository[T, R]) FindByFilterEager(predicate func(*gorm.DB) *gorm.DB, eagerAssociations []string) ([]T, error) {
	var result []T

	query := predicate(repository.Db)

	for _, association := range eagerAssociations {
		query = query.Preload(association)
	}

	if err := query.Find(&result).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}

	return result, nil
}

func (repository *GenericRepository[T, R]) FindOneByFilterEager(predicate func(*gorm.DB) *gorm.DB, eagerAssociations []string) (*T, error) {
	var entity T

	query := predicate(repository.Db)

	for _, association := range eagerAssociations {
		query = query.Preload(association)
	}

	if err := query.First(&entity).Error; err != nil {
		log.Panic("Error while fetching entity: ", err)
		return nil, err
	}

	return &entity, nil
}

func (repository *GenericRepository[T, R]) FindAllEager(eagerAssociations []string) ([]T, error) {
	var result []T

	query := repository.Db

	for _, association := range eagerAssociations {
		query = query.Preload(association)
	}

	if err := query.Find(&result).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}

	return result, nil

}

func (repository *GenericRepository[T, R]) FindByIdEager(id R, eagerAssociations []string) (T, error) {
	var entity T

	query := repository.Db

	for _, association := range eagerAssociations {
		query = query.Preload(association)
	}

	if err := query.First(&entity, id).Error; err != nil {
		log.Panic("Error while fetching entity: ", err)
		return entity, err
	}

	return entity, nil
}

func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}
