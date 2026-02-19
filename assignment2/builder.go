package assignment2

// builder interface

type BookingBuilder interface {
	SetUser(user User) BookingBuilder
	SetShow(show Show) BookingBuilder
	SetSeats(seats []Seat) BookingBuilder
	SetTotalAmount(amount float64) BookingBuilder
	AddSnack(snack Snack) BookingBuilder
	ApplyCoupon(code string) BookingBuilder
	SetGiftMessage(msg string) BookingBuilder
	SetRecipient(name string) BookingBuilder
	Build() Booking
}

type DefaultBookingBuilder struct {
	booking Booking
}

func NewBookingBuilder() BookingBuilder {
	return &DefaultBookingBuilder{
		booking: Booking{},
	}
}

func (b *DefaultBookingBuilder) SetUser(user User) BookingBuilder {
	b.booking.User = user
	return b
}

func (b *DefaultBookingBuilder) SetShow(show Show) BookingBuilder {
	b.booking.Show = show
	return b
}

func (b *DefaultBookingBuilder) SetSeats(seats []Seat) BookingBuilder {
	b.booking.Seats = seats
	return b
}

func (b *DefaultBookingBuilder)SetTotalAmount(total float64)BookingBuilder{
	b.booking.TotalAmount = total
	return b
}

func (b *DefaultBookingBuilder) AddSnack(snack Snack) BookingBuilder {
	b.booking.Snacks = append(b.booking.Snacks, snack)
	return b
}

func (b *DefaultBookingBuilder) ApplyCoupon(code string) BookingBuilder {
	b.booking.CouponCode = code
	return b
}

func (b *DefaultBookingBuilder) SetGiftMessage(msg string) BookingBuilder {
	b.booking.GiftMessage = msg
	return b
}

func (b *DefaultBookingBuilder) SetRecipient(name string) BookingBuilder {
	b.booking.Recipient = name
	return b
}

func (b *DefaultBookingBuilder) Build() Booking {
	return b.booking
}
