package eug

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repository[T schema.Tabler] struct {
	db *gorm.DB
	schema.Tabler
}

func NewRepository[T schema.Tabler](db *gorm.DB, entity schema.Tabler) Repository[T] {
	return Repository[T]{db, entity}
}

func (r *Repository[T]) First(ctx context.Context, model *T, filters ...Filter) error {
	return ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).First(&model).Error
}

func (r *Repository[T]) Find(ctx context.Context, models []T, filters ...Filter) error {
	return ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).Find(&models).Error
}

func (r *Repository[T]) Create(ctx context.Context, model *T) error {
	return r.db.Create(&model).Error
}

func (r *Repository[T]) Updates(ctx context.Context, model *T, updates *T) error {
	return r.db.Model(&model).Updates(&updates).Error
}

func (r *Repository[T]) Delete(ctx context.Context, model *T) error {
	return r.db.Delete(&model).Error
}
