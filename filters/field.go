package filters

import (
	"fmt"

	"github.com/shinhagunn/eug"
	"gorm.io/gorm"
)

func WithID(value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", value)
	}
}

func WithFieldEqual(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", field), value)
	}
}

func WithFieldNotEqual(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s <> ?", field), value)
	}
}

func WithFieldIn(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s IN ?", field), value)
	}
}

func WithFieldGreaterThan(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s > ?", field), value)
	}
}

func WithFieldLesserThan(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s < ?", field), value)
	}
}

func WithFieldGreaterOrEqual(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s >= ?", field), value)
	}
}

func WithFieldLesserOrEqual(field string, value interface{}) eug.Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s <= ?", field), value)
	}
}
