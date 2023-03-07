package main

import "discovery_language/internal/app/usecases"

func main() {
	travelledDistance := 27.8
	averageConsumption := 7.3

	usecases.DailyIntake(travelledDistance, averageConsumption)
}
