package usecases



func DailyIntake(travelledDistance float64, averageConsumption float64) float64 {
	dailyIntake := travelledDistance / averageConsumption
	return dailyIntake
}
