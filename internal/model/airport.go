package model

import "gorm.io/gorm"

type Airport struct {
	gorm.Model

	CityID   uint   `gorm:"not null;index"`
	Name     string `gorm:"not null"`
	IATACode string `gorm:"not null;uniqueIndex;size:3"`
	ICAOCode string `gorm:"not null;uniqueIndex;size:4"`

	City *City
}
