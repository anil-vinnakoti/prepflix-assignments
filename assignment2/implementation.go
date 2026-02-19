package assignment2

import (
	"errors"
	"fmt"
)

// booking service
type BookingService struct {
	allocator     SeatAllocationStrategy
	payment       PaymentGateway
	notification  NotificationService
	repo          BookingRepository
	ticketFactory TicketFactory
	builder       BookingBuilder
	logger        *Logger
}

func NewBookingService(
	allocator SeatAllocationStrategy,
	payment PaymentGateway,
	notification NotificationService,
	repo BookingRepository,
	ticketFactory TicketFactory,
	builder BookingBuilder,
	logger *Logger,
) *BookingService {
	return &BookingService{
		allocator:     allocator,
		payment:       payment,
		notification:  notification,
		repo:          repo,
		ticketFactory: ticketFactory,
		builder:       builder,
		logger:        logger,
	}
}

func (b *BookingService) Book(
	user User,
	show Show,
	seats []Seat,
) (*Booking, error) {

	if len(seats) == 0 {
		return nil, errors.New("no seats selected")
	}

	// 1️⃣ Calculate total price per seat
	total := 0.0

	for _, seat := range seats {
		ticket := b.ticketFactory.CreateTicket(seat.Type)
		if ticket == nil {
			return nil, errors.New("invalid seat type")
		}
		total += ticket.GetPrice()
	}

	// 2️⃣ Allocate seats
	if !b.allocator.AllocateSeats(show, seats) {
		return nil, errors.New("seats unavailable")
	}

	// 3️⃣ Process payment
	if err := b.payment.Pay(user, total); err != nil {
		b.allocator.ReleaseSeats(show, seats)
		return nil, fmt.Errorf("payment failed: %w", err)
	}

	// 4️⃣ Build booking
	booking := b.builder.
		SetUser(user).
		SetShow(show).
		SetSeats(seats).
		SetTotalAmount(total).
		Build()

	booking.Status = BookingConfirmed

	// 5️⃣ Save booking
	savedBooking, err := b.repo.SaveBooking(booking)
	if err != nil {
		b.allocator.ReleaseSeats(show, seats)
		return nil, err
	}

	// 6️⃣ Notify
	if err := b.notification.SendBookingConfirmation(user, savedBooking); err != nil {
		b.logger.Log("notification failed but booking confirmed")
	}

	b.logger.Log(fmt.Sprintf("Booking successful for user: %s", user.Name))

	return &savedBooking, nil
}
