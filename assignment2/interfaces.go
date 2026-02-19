package assignment2

type PricingStrategy interface {
	CalculatePrice(show Show, seat Seat, user User) float64
}

type PaymentGateway interface {
	Pay(user User, amount float64) error
}

type NotificationService interface {
	SendBookingConfirmation(user User, booking Booking) error
}

type SeatAllocationStrategy interface {
	AllocateSeats(show Show, seats []Seat) bool
	ReleaseSeats(show Show, seats []Seat)
}

type BookingRepository interface {
	SaveBooking(booking Booking) (Booking, error)
}
