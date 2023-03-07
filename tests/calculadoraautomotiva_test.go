package tests

import (
	"discovery_language/internal/app/entities"
	"discovery_language/internal/app/usecases"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Deve_Retornar_Calculo_de_Consumo_diario(t *testing.T) {
	car := entities.Car{
		Name:               "Siena",
		Model:              "ELX 2010",
		AverageConsumption: 7.3,
	}

	trip := entities.Trip{
		Name:              "Buscar Mozaum",
		TravelledDistance: 27.8,
		TripType:          "Diária a Trabalho",
	}

	fuel := entities.Fuel{
		Name:  "Etanol",
		Price: 3.89,
	}

	dailyIntake := usecases.DailyIntake(trip.TravelledDistance, car.AverageConsumption)
	dailyExpense := usecases.DailyExpense(dailyIntake, fuel.Price)

	// DailyExpense(dailyIntake, fuel.Price)

	result := math.Round(dailyExpense*100) / 100

	assert.Equal(t, 14.81, result, "Os numeros comparados não são iguais")
}
