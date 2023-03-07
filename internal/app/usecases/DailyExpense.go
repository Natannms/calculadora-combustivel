package usecases

type DailyExpenseUsecase struct {
	fuelPriceRepo FuelPriceRepository
}

func NewDailyExpenseUsecase(fuelPriceRepo FuelPriceRepository) *DailyExpenseUsecase {
	return &DailyExpenseUsecase{fuelPriceRepo: fuelPriceRepo}
}

func (uc *DailyExpenseUsecase) CalculateDailyExpense(dailyIntake float64, fuelPrice float64) (float64, error) {
	// Calcula o gasto diário
	dailyExpense := fuelPrice * dailyIntake

	return dailyExpense, nil
}
