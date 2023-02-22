package repository

import "gorm.io/gorm"

type Filter func(db *gorm.DB) *gorm.DB

func ApplyFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, f := range filters {
		f(db)
	}

	return db
}
