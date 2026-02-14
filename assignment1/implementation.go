package assignment1

import "errors"

// Pricing strategy instance
type BasePricingStrategy struct{}

func (b *BasePricingStrategy) CalculatePrice(show Show, seat Seat, user User) float64 {
	if seat.Type == "VIP" {
		return 200
	}
	if seat.Type == "Premium" {
		return 150
	}
	return 75
}

// PhonePe payment instance
type PhonePePaymentGateWay struct{}

func (p *PhonePePaymentGateWay) pay(user User, amount float64) error {
	// simulate payment processing
	return nil
}

// Email notification instance
type EmailNotificationService struct{}

func (e *EmailNotificationService) sendBookingConfirmation(user User, booking Booking) error {
	// simulate sending email
	return nil
}

// booking service
type BookingService struct {
	pricing      PricingStrategy
	allocator    SeatAllocationStrategy
	payment      PaymentGateway
	notification NotificationService
	repo         BookingRepository
}

func NewBookingService(
	pricing PricingStrategy,
	allocator SeatAllocationStrategy,
	payment PaymentGateway,
	notification NotificationService,
	repo BookingRepository,
) *BookingService {
	return &BookingService{
		pricing:      pricing,
		allocator:    allocator,
		payment:      payment,
		notification: notification,
		repo:         repo,
	}
}

func (b *BookingService) Book(user User, show Show, seats []Seat) (*Booking, error) {
	// calculate price
	total := 0.0
	for _, seat := range seats {
		total += b.pricing.CalculatePrice(show, seat, user)
	}

	// allocate seats
	if !b.allocator.AllocateSeat(show, seats) {
		return nil, errors.New("seats unavailable")
	}

	// charge payment
	if err := b.payment.pay(user, total); err != nil {
		b.allocator.ReleaseSeats(show, seats) // releasing seats as payment failed
		return nil, errors.New("payment failed")
	}

	booking := Booking{
		// ID: uuid.New.String(),
		User:   user,
		Show:   show,
		Seats:  seats,
		Amount: total,
		Status: BookingConfirmed,
	}

	// save booking in repository
	savedBooking, err := b.repo.SaveBooking(booking)
	if err != nil {
		return nil, err
	}

	// send booking confirmation
	_ = b.notification.sendBookingConfirmation(user, savedBooking)

	return &savedBooking, nil
}
