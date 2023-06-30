package eug

import "gorm.io/gorm"

func ApplyFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, filter := range filters {
		filter(db)
	}

	return db
}
