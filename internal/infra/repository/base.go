package repository

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *BaseRepository[T]) FindByID(id uint, entity *T) error {
	return r.DB.First(entity, id).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}

func (r *BaseRepository[T]) List(entities *[]T, conditions map[string]interface{}) error {
	return r.DB.Where(conditions).Find(entities).Error
}

func (r *BaseRepository[T]) FindByCondition(condition interface{}, args ...interface{}) ([]T, error) {
	var entities []T
	err := r.DB.Where(condition, args...).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepository[T]) FindOneByCondition(condition interface{}, args ...interface{}) (*T, error) {
	var entity T
	err := r.DB.Where(condition, args...).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
