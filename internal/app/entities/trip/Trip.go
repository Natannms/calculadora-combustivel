package trip

type Trip struct {
	Name              string
	TravelledDistance float64
	TripType          string
}

func (t Trip) GetTravelledDistance() float64 {
	return t.TravelledDistance
}
