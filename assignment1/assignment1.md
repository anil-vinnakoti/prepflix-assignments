1. UML class diagram:
   User 1.._ ------ Booking
   Booking ------ 1 Show
   Show ------ 1 Movie
   Show ------ 1 Screen
   Screen 1.._ ------ Seat

2. Class Responsibility Table
   Component Responsibility
   Movie Stores movie metadata
   Theater Theater information
   Screen Contains seats
   Seat Seat information & type
   User User information
   Show Represents a screening
   Booking Booking data
   BookingService controls booking workflow
   PricingStrategy Calculates ticket price
   SeatAllocationStrategy Manages seat allocation & de-allocation
   PaymentGateway Processes payments
   NotificationService Sends booking confirmation
   BookingRepository Persists booking data

3. Each SOLID principle explanation
   1. Single Responsibility :- A class should have only one reason to change.
    Entities contain only data.
    Business logic exists in service classes.
    Seat allocation logic exists only in SeatAllocationStrategy.
    Pricing logic exists only in PricingStrategy.
    
    EXAMPLE:
    If pricing rules change → modify PricingStrategy.
    If payment provider changes → modify PaymentGateway implementation.

   2. Open/Closed :- Entities should be open for extension, closed for modification.
    Instead of modifying BookingService when pricing rules change, we introduce new implementations of PricingStrategy.
    EXAMPLE:- SurgePricingStrategy
    BookingService remains unchanged. This makes the system extensible without modifying stable code

   3. Liskov Substitution :- Subtypes must be substitutable for their base types.
    A pricing strategy must always return a valid price.
    A payment gateway must return clear success or failure.
    If an implementation breaks expected behavior, it violates LSP.

   4. Interface Segregation :- Clients should not depend on methods they do not use.
    Instead of one large interface, i have defined PricingStrategy, SeatAllocationStrategy, PaymentGateway, NotificationService.
    Each interface has a single responsibility, easy to mock, keeps implementations simple and, this avoids “fat interfaces.”

   5. Dependency Inversion :- High-level modules should depend on abstractions, not concrete implementations.
    BookingService depends only on PricingStrategy, SeatAllocationStrategy, PaymentGateway, NotificationService, BookingRepository interfaces
    It does NOT depend on GPay, Paytm, Email provider, Specific pricing rule and Concrete implementations are injected during initialization.
    With this the code is loosely coupled and easier to test with mocks
