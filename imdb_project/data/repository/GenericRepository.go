package repository

import (
	"gorm.io/gorm"
	"log"
)

type IGenericRepository[T any, R any] interface {
	FindAll() ([]T, error)
	FindByID(id R) (T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	DeleteById(id R) (bool, error)
	FindByFilter(predicate func(*gorm.DB) *gorm.DB) ([]T, error)
	FindOneByFilter(predicate func(*gorm.DB) *gorm.DB) (*T, error)
}

type GenericRepository[T, R any] struct {
	Db *gorm.DB
}

func NewGenericRepository[T, R any](db *gorm.DB) IGenericRepository[T, R] {
	return &GenericRepository[T, R]{Db: db}
}
func (r *GenericRepository[T, R]) FindAll() ([]T, error) {
	var entities []T
	if err := r.Db.Find(&entities).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}
	return entities, nil
}

func (r *GenericRepository[T, R]) FindByID(id R) (T, error) {
	var entity T
	if err := r.Db.First(&entity, id).Error; err != nil {
		log.Panic("Error while fetching entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (r *GenericRepository[T, R]) Create(entity *T) (*T, error) {
	if err := r.Db.Create(entity).Error; err != nil {
		log.Panic("Error while creating entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (r *GenericRepository[T, R]) Update(entity *T) (*T, error) {
	if err := r.Db.Save(entity).Error; err != nil {
		log.Panic("Error while updating entity: ", err)
		return entity, err
	}
	return entity, nil
}

func (r *GenericRepository[T, R]) DeleteById(id R) (bool, error) {
	if err := r.Db.Delete(new(T), id).Error; err != nil {
		log.Panic("Error while deleting entity: ", err)
		return false, err
	}
	return true, nil
}

func (r *GenericRepository[T, R]) FindByFilter(predicate func(*gorm.DB) *gorm.DB) ([]T, error) {
	var entities []T
	query := predicate(r.Db)
	if err := query.Find(&entities).Error; err != nil {
		log.Panic("Error while fetching entities: ", err)
		return nil, err
	}
	return entities, nil
}

func (r *GenericRepository[T, R]) FindOneByFilter(predicate func(*gorm.DB) *gorm.DB) (*T, error) {
	var entity T
	query := predicate(r.Db)
	if err := query.First(&entity).Error; err != nil {
		log.Panic("Error while fetching entity: ", err)
		return nil, err
	}
	return &entity, nil
}
