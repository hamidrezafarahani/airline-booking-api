package model

import "gorm.io/gorm"

type TicketStatus string

const (
	TicketStatusIssued    TicketStatus = "issued"
	TicketStatusRefunded  TicketStatus = "refunded"
	TicketStatusCancelled TicketStatus = "cancelled"
)

type Ticket struct {
	gorm.Model

	TicketNumber string `gorm:"not null;uniqueIndex"`

	FlightID    uint `gorm:"not null;index"`
	PassengerID uint `gorm:"not null;index"`
	AgencyID    uint `gorm:"not null;index"`

	Price int64 `gorm:"not null"`

	Status TicketStatus `gorm:"not null;index"`

	Flight    *Flight
	Passenger *Passenger
	Agency    *Agency
}
