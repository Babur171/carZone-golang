package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()

	yearInt, _ := strconv.Atoi(year)

	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year muct be between 1886 and current year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	validetFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, valueType := range validetFuelTypes {
		if fuelType == valueType {
			return nil
		}
	}
	return errors.New("fueltype must be Petrol, Diesel, Electric and hybrid")

}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("engineid is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than 0")
	}
	if engine.CarRange <= 0 {
		return errors.New("CarRange must be greater than 0")
	}
	return nil

}

func validatePrice(price int64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil

}
