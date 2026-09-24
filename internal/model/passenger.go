package model

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	NationalID string `gorm:"not null;uniqueIndex"`

	Tickets []Ticket
}
