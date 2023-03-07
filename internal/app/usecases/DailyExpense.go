package usecases

type DailyExpenseRepository interface {
	CalculateDailyExpense(vehicleID int, fuelPrice float64, distance float64) float64
}

type DailyExpenseUseCase struct {
	dailyExpenseRepo DailyExpenseRepository
}

func NewDailyExpenseCalculation(dailyExpenseRepo DailyExpenseRepository) *DailyExpenseUseCase {
	return &DailyExpenseUseCase{dailyExpenseRepo: dailyExpenseRepo}
}

func (de *DailyExpenseUseCase) DailyExpense(dailyIntake float64, fuelPrice float64) float64 {
	result := dailyIntake * fuelPrice
	return result
}
