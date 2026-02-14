## UML Class Diagram

User 1..* ------ Booking

Booking ------ 1 Show

Show ------ 1 Movie

Show ------ 1 Screen

Screen 1..* ------ Seat

---

## Class Responsibility Table

| Component | Responsibility |
|------------|----------------|
| Movie | Stores movie metadata |
| Theater | Theater information |
| Screen | Contains seats |
| Seat | Seat information & type |
| User | User information |
| Show | Represents a screening |
| Booking | Booking data |
| BookingService | Controls booking workflow |
| PricingStrategy | Calculates ticket price |
| SeatAllocationStrategy | Manages seat allocation & de-allocation |
| PaymentGateway | Processes payments |
| NotificationService | Sends booking confirmation |
| BookingRepository | Persists booking data |

---

## SOLID Principles Explanation

---

### 1. Single Responsibility Principle (SRP)

> A class should have only one reason to change.

- Entities contain only data.
- Business logic exists in service classes.
- Seat allocation logic exists only in `SeatAllocationStrategy`.
- Pricing logic exists only in `PricingStrategy`.

**Example:**

- If pricing rules change → modify `PricingStrategy`.
- If payment provider changes → modify `PaymentGateway` implementation.

---

### 2. Open/Closed Principle (OCP)

> Software entities should be open for extension but closed for modification.

Instead of modifying `BookingService` when pricing rules change,  
we introduce new implementations of `PricingStrategy`.

**Example:**

- `SurgePricingStrategy`

`BookingService` remains unchanged.  
This makes the system extensible without modifying stable code.

---

### 3. Liskov Substitution Principle (LSP)

> Subtypes must be substitutable for their base types.

- A pricing strategy must always return a valid price.
- A payment gateway must return clear success or failure.
- If an implementation breaks expected behavior, it violates LSP.

---

### 4. Interface Segregation Principle (ISP)

> Clients should not depend on methods they do not use.

Instead of one large interface, the system defines:

- `PricingStrategy`
- `SeatAllocationStrategy`
- `PaymentGateway`
- `NotificationService`

Each interface:
- Has a single responsibility
- Is easy to mock
- Keeps implementations simple
- Avoids "fat interfaces"

---

### 5. Dependency Inversion Principle (DIP)

> High-level modules should depend on abstractions, not concrete implementations.

`BookingService` depends only on:

- `PricingStrategy`
- `SeatAllocationStrategy`
- `PaymentGateway`
- `NotificationService`
- `BookingRepository`

It does **NOT** depend on:

- GPay
- Paytm
- Email provider
- Specific pricing rule

Concrete implementations are injected during initialization.

This makes the code:
- Loosely coupled
- Easier to test
- More maintainable
