package model

import "gorm.io/gorm"

type Country struct {
	gorm.Model

	Name string `gorm:"not null;uniqueIndex"`
	Code string `gorm:"not null;uniqueIndex;size:3"`

	Provinces []Province
}
