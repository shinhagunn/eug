package filters

import (
	"fmt"

	"github.com/shinhagunn/eug"
	"gorm.io/gorm"
)

func OrderByAsc(field string) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s asc", field))
	}
}

func OrderByDesc(field string) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s desc", field))
	}
}

func WithLimit(limit int) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func WithOffset(offset int) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
