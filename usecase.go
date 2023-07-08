package eug

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Usecase[T schema.Tabler] struct {
	repo Repository[T]
}

type IUsecase[T schema.Tabler] interface {
	First(ctx context.Context, filters ...Filter) (*T, error)
	Find(ctx context.Context, filters ...Filter) []T
	Create(ctx context.Context, model *T) error
	Updates(ctx context.Context, model *T, updates *T) error
	Delete(ctx context.Context, model *T) error
}

func (u Usecase[T]) First(ctx context.Context, filters ...Filter) (*T, error) {
	model := new(T)

	err := u.repo.First(ctx, model, filters...)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (u Usecase[T]) Find(ctx context.Context, filters ...Filter) []T {
	models := []T{}

	if err := u.repo.Find(ctx, models, filters...); err != nil {
		return nil
	}

	return models
}

func (u Usecase[T]) Create(ctx context.Context, model *T) error {
	return u.repo.Create(ctx, model)
}

func (u Usecase[T]) Updates(ctx context.Context, model *T, updates *T) error {
	return u.repo.Updates(ctx, model, updates)
}

func (u Usecase[T]) Delete(ctx context.Context, model *T) error {
	return u.repo.Delete(ctx, model)
}
