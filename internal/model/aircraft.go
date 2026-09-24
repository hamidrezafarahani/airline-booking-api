package model

import "gorm.io/gorm"

type Aircraft struct {
	gorm.Model

	AircraftTypeID     uint   `gorm:"not null;index"`
	RegistrationNumber string `gorm:"not null;uniqueIndex"`

	AircraftType *AircraftType

	Flights []Flight
}
