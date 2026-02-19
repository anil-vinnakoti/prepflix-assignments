package assignment2

import "time"

type Movie struct {
	ID       string
	Title    string
	Language string
	Rating   int
	Duration time.Duration
}

type Theater struct {
	ID      string
	Name    string
	Address string
	Screens int
}

type Screen struct {
	ID    string
	Name  string
	Seats []int
}

type Seat struct {
	ID     string
	Type   string // VIP, Premium, Economy
	Row    string
	Number int
}

type User struct {
	ID    string
	Name  string
	Email string
}

type Show struct {
	ID       string
	Movie    Movie
	Screen   Screen
	ShowTime time.Time
}

type Snack struct{
	Name string
	Price float64
}

type BookingStatus string

const (
	BookingConfirmed BookingStatus = "CONFIRMED"
	BookingFailed    BookingStatus = "FAILED"
)


type Booking struct{
	User User
	Show Show
	Seats []Seat
	Snacks []Snack
	CouponCode string
	GiftMessage string
	Recipient string
	TotalAmount float64
	Status BookingStatus
}

