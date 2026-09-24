package postgres

import (
	"fmt"
	"time"

	"github.com/hamidrezafarahani/airline-booking-api/internal/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Country
		var iran model.Country

		if err := tx.
			Where("code = ?", "IR").
			FirstOrCreate(
				&iran,
				model.Country{
					Name: "Iran",
					Code: "IR",
				},
			).Error; err != nil {
			return fmt.Errorf("seed country: %w", err)
		}

		// Province
		var tehranProvince model.Province

		if err := tx.
			Where(
				"country_id = ? AND name = ?",
				iran.ID,
				"Tehran",
			).
			FirstOrCreate(
				&tehranProvince,
				model.Province{
					CountryID: iran.ID,
					Name:      "Tehran",
				},
			).Error; err != nil {
			return fmt.Errorf("seed province: %w", err)
		}

		// City
		var tehranCity model.City

		if err := tx.
			Where(
				"province_id = ? AND name = ?",
				tehranProvince.ID,
				"Tehran",
			).
			FirstOrCreate(
				&tehranCity,
				model.City{
					ProvinceID: tehranProvince.ID,
					Name:       "Tehran",
				},
			).Error; err != nil {
			return fmt.Errorf("seed city: %w", err)
		}

		// Airport
		var ika model.Airport

		if err := tx.
			Where("iata_code = ?", "IKA").
			FirstOrCreate(
				&ika,
				model.Airport{
					CityID:   tehranCity.ID,
					Name:     "Imam Khomeini International Airport",
					IATACode: "IKA",
					ICAOCode: "OIIE",
				},
			).Error; err != nil {
			return fmt.Errorf("seed airport: %w", err)
		}

		// Aircraft Type
		var a320 model.AircraftType

		if err := tx.
			Where("name = ?", "Airbus A320").
			FirstOrCreate(
				&a320,
				model.AircraftType{
					Name:         "Airbus A320",
					Manufacturer: "Airbus",
					Capacity:     180,
				},
			).Error; err != nil {
			return fmt.Errorf("seed aircraft type: %w", err)
		}

		// Aircraft
		var aircraft model.Aircraft

		if err := tx.
			Where(
				"registration_number = ?",
				"EP-ABC",
			).
			FirstOrCreate(
				&aircraft,
				model.Aircraft{
					AircraftTypeID:     a320.ID,
					RegistrationNumber: "EP-ABC",
				},
			).Error; err != nil {
			return fmt.Errorf("seed aircraft: %w", err)
		}

		// Agency
		var agency model.Agency

		if err := tx.
			Where("code = ?", "TEST").
			FirstOrCreate(
				&agency,
				model.Agency{
					Name:   "Test Travel Agency",
					Code:   "TEST",
					Status: model.AgencyStatusActive,
				},
			).Error; err != nil {
			return fmt.Errorf("seed agency: %w", err)
		}

		// Wallet
		var wallet model.Wallet

		if err := tx.
			Where("agency_id = ?", agency.ID).
			FirstOrCreate(
				&wallet,
				model.Wallet{
					AgencyID: agency.ID,
					Balance:  1_000_000_000,
				},
			).Error; err != nil {
			return fmt.Errorf("seed wallet: %w", err)
		}

		// Passenger
		var passenger model.Passenger

		if err := tx.
			Where("national_id = ?", "0012345678").
			FirstOrCreate(
				&passenger,
				model.Passenger{
					FirstName:  "Ali",
					LastName:   "Ahmadi",
					NationalID: "0012345678",
				},
			).Error; err != nil {
			return fmt.Errorf("seed passenger: %w", err)
		}

		// Flight
		var flight model.Flight

		departureAt := time.Now().
			Add(24 * time.Hour).
			Truncate(time.Second)

		arrivalAt := departureAt.Add(3 * time.Hour)

		if err := tx.
			Where("flight_number = ?", "IR100").
			First(&flight).Error; err != nil {

			if err != gorm.ErrRecordNotFound {
				return fmt.Errorf("find flight: %w", err)
			}

			flight = model.Flight{
				FlightNumber:       "IR100",
				AircraftID:         aircraft.ID,
				DepartureAirportID: ika.ID,
				ArrivalAirportID:   ika.ID,
				DepartureAt:        departureAt,
				ArrivalAt:          arrivalAt,
				BasePrice:          50_000_000,
				Status:             model.FlightStatusScheduled,
			}

			if err := tx.Create(&flight).Error; err != nil {
				return fmt.Errorf("seed flight: %w", err)
			}
		}

		return nil
	})
}
