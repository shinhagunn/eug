package repositories

import (
	"context"

	"github.com/shinhagunn/eug"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type repository[T schema.Tabler] struct {
	db *gorm.DB
	schema.Tabler
}

func NewRepository[T schema.Tabler](db *gorm.DB, entity schema.Tabler) *repository[T] {
	return &repository[T]{db, entity}
}

func (r *repository[T]) First(ctx context.Context, model *T, filters ...eug.Filter) error {
	return eug.ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).First(&model).Error
}

func (r *repository[T]) Find(ctx context.Context, models []T, filters ...eug.Filter) error {
	return eug.ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).Find(&models).Error
}

func (r *repository[T]) Create(ctx context.Context, model *T) error {
	return r.db.Create(&model).Error
}

func (r *repository[T]) Updates(ctx context.Context, model *T, updates *T) error {
	return r.db.Model(&model).Updates(&updates).Error
}

func (r *repository[T]) Delete(ctx context.Context, model *T) error {
	return r.db.Delete(&model).Error
}
