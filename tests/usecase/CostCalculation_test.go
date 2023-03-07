package tests

import (
	"discovery_language/internal/app/entities/car"
	"discovery_language/internal/app/entities/fuel"
	"discovery_language/internal/app/entities/trip"
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

	trip := trip.Trip{
		Name:              "Bairro ao lado",
		TravelledDistance: 27.8,
		TripType:          "Daily",
	}

	car := car.Car{
		Name:               "Siena",
		Model:              "ELX 2010",
		AverageConsumption: 7.3,
	}

	fuel := fuel.Fuel{
		Name:  "Etanol",
		Price: 4.97,
	}

	expectedCost := 4.97 * (trip.GetTravelledDistance() / car.GetAverageConsumption())

	cost, err := cc.CalculateCostByFuelType(trip.GetTravelledDistance(), car.AverageConsumption, fuel.GetPrice())
	if err != nil {
		t.Errorf("Error calculating cost %v", err)
	}

	assert.Equal(t, expectedCost, cost, "VALOR ESPERADO É DIFERENTE DO VALOR RECEBIDO")
}
