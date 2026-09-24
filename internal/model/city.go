package model

import "gorm.io/gorm"

type City struct {
	gorm.Model

	ProvinceID uint   `gorm:"not null;index;uniqueIndex:idx_city_province_name"`
	Name       string `gorm:"not null;uniqueIndex:idx_city_province_name"`

	Province *Province
	Airports []Airport
}
