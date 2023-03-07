package usecases

// type DailyExpenseUseCase interface {
// 	DailyExpense(dailyIntake float64, fuelPrice float64) float64
// }

// type DailyExpenseRepository interface {
// 	GenerateDailyExpense(dailyIntake float64, fuelPrice float64) float64
// }

// type DailyExpenseInteractor struct {
// 	DailyExpenseRepository DailyExpenseRepository
// }

// func (i *DailyExpenseInteractor) DailyExpense(dailyIntake float64, fuelPrice float64) float64 {

// 	return 1
// }

func DailyExpense(dailyIntake float64, fuelPrice float64) float64 {
	result := dailyIntake * fuelPrice
	return result
}
