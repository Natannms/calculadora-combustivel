package car

type Car struct {
	Name               string
	Model              string
	AverageConsumption float64
}

func (c Car) GetAverageConsumption() float64 {
	return c.AverageConsumption
}
