package filters

import (
	"fmt"

	"github.com/shinhagunn/eug/repository"
	"gorm.io/gorm"
)

func WithFieldEqual(field string, value interface{}) repository.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", field), value)
	}
}
