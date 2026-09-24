package postgres

import (
	"github.com/hamidrezafarahani/airline-booking-api/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Country{},
		&model.Province{},
		&model.City{},
		&model.Airport{},

		&model.AircraftType{},
		&model.Aircraft{},

		&model.Flight{},

		&model.Agency{},
		&model.APIKey{},
		&model.Wallet{},
		&model.WalletTransaction{},

		&model.Passenger{},
		&model.Ticket{},
	)

}
