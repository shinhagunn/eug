package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type repository[T schema.Tabler] struct {
	*gorm.DB
	schema.Tabler
}

type Repository[T schema.Tabler] interface {
	First(ctx context.Context, model interface{}, filters ...Filter) error
	Find(ctx context.Context, models interface{}, filters ...Filter) error
	Create(ctx context.Context, models interface{}, filters ...Filter) error
	Updates(ctx context.Context, model interface{}, value interface{}, filters ...Filter) error
	Delete(ctx context.Context, model interface{}, filters ...Filter) error
}

func NewRepository[T schema.Tabler](db *gorm.DB, entity T) Repository[T] {
	return &repository[T]{db, entity}
}

func (r repository[T]) First(ctx context.Context, model interface{}, filters ...Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).First(model).Error
}

func (r repository[T]) Find(ctx context.Context, models interface{}, filters ...Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Find(models).Error
}

func (r repository[T]) Create(ctx context.Context, models interface{}, filters ...Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Create(models).Error
}

func (r repository[T]) Updates(ctx context.Context, model interface{}, value interface{}, filters ...Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Model(model), filters).Updates(value).Error
}

func (r repository[T]) Delete(ctx context.Context, model interface{}, filters ...Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Delete(model).Error
}
