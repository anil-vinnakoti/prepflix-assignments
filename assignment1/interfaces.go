package assignment1

type PricingStrategy interface {
	CalculatePrice(show Show, seat Seat, user User) float64
}

type SeatAllocationStrategy interface {
	AllocateSeat(show Show, seats []Seat) bool
	ReleaseSeats(show Show, seats []Seat)
}

type PaymentGateway interface {
	pay(user User, amount float64) error
}

type NotificationService interface {
	sendBookingConfirmation(user User, booking Booking) error
}

type BookingRepository interface {
	SaveBooking(booking Booking) (Booking, error)
}
