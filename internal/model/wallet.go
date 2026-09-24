package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model

	AgencyID uint `gorm:"not null;uniqueIndex"`

	Balance int64 `gorm:"not null;default:0"`

	Agency *Agency

	Transactions []WalletTransaction
}
