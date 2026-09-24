package model

import "gorm.io/gorm"

type AircraftType struct {
	gorm.Model

	Name         string `gorm:"not null;uniqueIndex"`
	Manufacturer string `gorm:"not null"`
	Capacity     int    `gorm:"not null"`

	Aircrafts []Aircraft
}
