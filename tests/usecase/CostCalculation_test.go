package tests

import (
	"discovery_language/internal/app/entities"
	"discovery_language/internal/app/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockFuelPriceRepository struct{}

func (m *mockFuelPriceRepository) GetPrice(fuelType string) (float64, error) {
	if fuelType == "gasolina" {
		return 4.50, nil
	}
	return 0, errors.New("Fuel type not found")
}

func TestCostCalculationUsecase_CalculateCostByFuelType(t *testing.T) {
	cc := usecases.NewCostCalculationUsecase(&mockFuelPriceRepository{})

	trip := entities.Trip{
		Name:              "Bairro ao lado",
		TravelledDistance: 27.8,
		TripType:          "Daily",
	}

	car := entities.Car{
		Name:               "Siena",
		Model:              "ELX 2010",
		AverageConsumption: 7.3,
	}

	fuel := entities.Fuel{
		Name:  "Etanol",
		Price: 4.97,
	}

	expectedCost := 4.97 * (trip.TravelledDistance / car.AverageConsumption)

	cost, err := cc.CalculateCostByFuelType(trip.TravelledDistance, car.AverageConsumption, fuel.Price)
	if err != nil {
		t.Errorf("Error calculating cost %v", err)
	}

	assert.Equal(t, expectedCost, cost, "VALOR ESPERADO É DIFERENTE DO VALOR RECEBIDO")
}
