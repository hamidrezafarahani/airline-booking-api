package model

import "gorm.io/gorm"

type WalletTransactionType string

const (
	WalletTransactionDebit  WalletTransactionType = "debit"
	WalletTransactionCredit WalletTransactionType = "credit"
)

type WalletTransaction struct {
	gorm.Model

	WalletID uint `gorm:"not null;index"`

	Type WalletTransactionType `gorm:"not null"`

	Amount int64 `gorm:"not null"`

	BalanceBefore int64 `gorm:"not null"`
	BalanceAfter  int64 `gorm:"not null"`

	ReferenceType string `gorm:"size:50"`
	ReferenceID   *uint

	Wallet *Wallet
}
