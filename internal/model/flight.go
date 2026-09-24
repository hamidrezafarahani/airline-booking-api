package model

import (
	"time"

	"gorm.io/gorm"
)

type FlightStatus string

const (
	FlightStatusScheduled FlightStatus = "scheduled"
	FlightStatusCancelled FlightStatus = "cancelled"
	FlightStatusCompleted FlightStatus = "completed"
)

type Flight struct {
	gorm.Model

	FlightNumber string `gorm:"not null;index"`

	AircraftID uint `gorm:"not null;index"`

	DepartureAirportID uint `gorm:"not null;index"`
	ArrivalAirportID   uint `gorm:"not null;index"`

	DepartureAt time.Time `gorm:"not null;index"`
	ArrivalAt   time.Time `gorm:"not null"`

	BasePrice int64 `gorm:"not null"`

	Status FlightStatus `gorm:"not null;index"`

	Aircraft *Aircraft

	DepartureAirport *Airport `gorm:"foreignKey:DepartureAirportID"`
	ArrivalAirport   *Airport `gorm:"foreignKey:ArrivalAirportID"`

	Tickets []Ticket
}
