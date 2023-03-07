package usecases

type FuelPriceRepository interface {
	GetPrice(fuelType string) (float64, error)
}

type CostCalculationUsecase struct {
	fuelPriceRepo FuelPriceRepository
}

func NewCostCalculationUsecase(fuelPriceRepo FuelPriceRepository) *CostCalculationUsecase {
	return &CostCalculationUsecase{fuelPriceRepo: fuelPriceRepo}
}

func (cc *CostCalculationUsecase) CalculateCost(tripDistance, fuelConsumption, fuelPrice float64) (float64, error) {
	totalFuelCost := fuelPrice * (tripDistance / fuelConsumption)
	return totalFuelCost, nil
}

func (cc *CostCalculationUsecase) CalculateCostByFuelType(tripDistance, fuelConsumption float64, fuelType string) (float64, error) {
	fuelPrice, err := cc.fuelPriceRepo.GetPrice(fuelType)
	if err != nil {
		return 0, err
	}

	totalFuelCost := fuelPrice * (tripDistance / fuelConsumption)
	return totalFuelCost, nil
}

// func DailyIntake(travelledDistance float64, averageConsumption float64) float64 {
// 	dailyIntake := travelledDistance / averageConsumption
// 	return dailyIntake
// }
