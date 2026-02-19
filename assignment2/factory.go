package assignment2

// ticket interface
type Ticket interface {
	GetPrice() float64
	GetAmenities() []string
	GetType() string
}

// standard ticket concrete class(struct)
type StandarTicket struct{}

func (s *StandarTicket) GetPrice() float64 {
	return 150
}
func (s *StandarTicket) GetAmenities() []string {
	return []string{"basic seat"}
}
func (s *StandarTicket) GetType() string {
	return "STANDARD"
}

// premium ticket concrete class(struct)
type PremiumTicket struct{}

func (p *PremiumTicket) GetPrice() float64 {
	return 250
}
func (p *PremiumTicket) GetAmenities() []string {
	return []string{"extra legroom"}
}
func (p *PremiumTicket) GetType() string {
	return "PREMIUM"
}

// IMAX ticket concrete class(struct)
type IMAXTicket struct{}

func (i *IMAXTicket) GetPrice() float64 {
	return 435
}
func (i *IMAXTicket) GetAmenities() []string {
	return []string{"IMAX screen", "dolby audio"}
}
func (i *IMAXTicket) GetType() string {
	return "IMAX"
}

// recliner ticket concrete class(struct)
type ReclinerTicket struct{}

func (r *ReclinerTicket) GetPrice() float64 {
	return 300
}
func (r *ReclinerTicket) GetAmenities() []string {
	return []string{"recliner seat"}
}
func (r *ReclinerTicket) GetType() string {
	return "RECLINER"
}

// factory interface
type TicketFactory interface {
	CreateTicket(ticketType string) Ticket
}

type DefaultTicketFactory struct{}

func (D *DefaultTicketFactory) CreateTicket(ticketType string) Ticket {
	switch ticketType {
	case "PREMIUM":
		return &PremiumTicket{}
	case "IMAX":
		return &IMAXTicket{}
	case "RECLINER":
		return &ReclinerTicket{}
	case "GIFT": // extension scenario
		return &GiftTicket{}
	default:
		return &StandarTicket{}
	}
}

// OCP extension - GiftTicket
type GiftTicket struct{}

func (g *GiftTicket) GetPrice() float64 {
	return 0
}
func (g *GiftTicket) GetAmenities() []string {
	return []string{"gift voucher"}
}
func (g *GiftTicket) GetType() string {
	return "GIFT"
}
