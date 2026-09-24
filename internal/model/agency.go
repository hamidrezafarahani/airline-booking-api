package model

import "gorm.io/gorm"

type AgencyStatus string

const (
	AgencyStatusActive    AgencyStatus = "active"
	AgencyStatusSuspended AgencyStatus = "suspended"
)

type Agency struct {
	gorm.Model

	Name   string       `gorm:"not null"`
	Code   string       `gorm:"not null;uniqueIndex"`
	Status AgencyStatus `gorm:"not null;index"`

	APIKeys []APIKey

	Wallet *Wallet

	Tickets []Ticket
}
