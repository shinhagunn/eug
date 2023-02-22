package filters

import (
	"github.com/shinhagunn/eug/repository"
	"gorm.io/gorm"
)

func WithOmit(fields ...string) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Omit(fields...)
	}
}
