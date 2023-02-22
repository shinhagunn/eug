package filters

import (
	"github.com/shinhagunn/eug/repository"
	"gorm.io/gorm"
)

func WithID(value interface{}) repository.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", value)
	}
}
