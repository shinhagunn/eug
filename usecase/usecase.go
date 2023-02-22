package usecase

import (
	"context"
	"errors"

	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/eug/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Usecase[T schema.Tabler] struct {
	Repository repository.Repository[T]
	Omits      []string
}

type IUsecase[T schema.Tabler] interface {
	First(ctx context.Context, filters ...repository.Filter) (*T, error)
	Find(ctx context.Context, filters ...repository.Filter) []*T
	Create(ctx context.Context, model *T, fs ...repository.Filter)
	Updates(ctx context.Context, model *T, value interface{}, fs ...repository.Filter)
	Delete(ctx context.Context, model *T, fs ...repository.Filter)
}

func (u Usecase[T]) First(ctx context.Context, filters ...repository.Filter) (*T, error) {
	var value *T
	if err := u.Repository.First(ctx, &value, filters...); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return value, nil
}

func (u Usecase[T]) Find(ctx context.Context, filters ...repository.Filter) []*T {
	var values []*T
	if err := u.Repository.Find(ctx, &values, filters...); err != nil {
		panic(err)
	}

	return values
}

func (u Usecase[T]) Create(ctx context.Context, model *T, fs ...repository.Filter) {
	fs = append(fs, filters.WithOmit(u.Omits...))

	if err := u.Repository.Create(ctx, model, fs...); err != nil {
		panic(err)
	}
}

func (u Usecase[T]) Updates(ctx context.Context, model *T, value interface{}, fs ...repository.Filter) {
	fs = append(fs, filters.WithOmit(u.Omits...))

	if err := u.Repository.Updates(ctx, model, value, fs...); err != nil {
		panic(err)
	}
}

func (u Usecase[T]) Delete(ctx context.Context, model *T, fs ...repository.Filter) {
	fs = append(fs, filters.WithOmit(u.Omits...))

	if err := u.Repository.Delete(ctx, model, fs...); err != nil {
		panic(err)
	}
}
