package assignment1

type SurgePricingStartegy struct {
	Base PricingStrategy
}

func (s *SurgePricingStartegy) CalculatePrice(show Show, seat Seat, user User) float64 {
	basePrice := s.CalculatePrice(show, seat, user)

	// weekend surge
	if show.ShowTime.Weekday().String() == "Saturday" || show.ShowTime.Weekday().String() == "Sunday" {
		return basePrice * 1.3
	}

	return basePrice
}
