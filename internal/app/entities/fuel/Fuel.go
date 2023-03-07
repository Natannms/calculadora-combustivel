package fuel

type Fuel struct {
	Name  string
	Price float64
}

func (f Fuel) GetPrice() float64 {
	return f.Price
}
