package model

import "gorm.io/gorm"

type Province struct {
	gorm.Model

	CountryID uint   `gorm:"not null;index;uniqueIndex:idx_province_country_name"`
	Name      string `gorm:"not null;uniqueIndex:idx_province_country_name"`

	Country *Country
	Cities  []City
}
