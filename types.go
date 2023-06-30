package eug

import "gorm.io/gorm"

type Filter func(*gorm.DB) *gorm.DB
