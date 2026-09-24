package model

import (
	"time"

	"gorm.io/gorm"
)

type APIKey struct {
	gorm.Model

	AgencyID uint `gorm:"not null;index"`

	Name string `gorm:"not null"`

	KeyHash string `gorm:"not null;uniqueIndex"`

	ExpiresAt *time.Time
	RevokedAt *time.Time

	Agency *Agency
}
