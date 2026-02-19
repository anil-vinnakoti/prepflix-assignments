![alt text](image.png)

# System Flow Explanation

### 1. Price Calculation

- Each seat contains a `Type`.
- For every seat:
  - Factory creates corresponding `Ticket`.
  - `Ticket.GetPrice()` returns seat price.
- Total is computed by summing prices of all seats.
- Supports mixed seat types (VIP + Premium + Standard).

---

### 2. Booking Construction

- `BookingBuilder` constructs Booking object.
- Ensures controlled object creation.
- Avoids large constructors.
- Improves readability and maintainability.

---

### 3. Logging

- Logger is a Singleton.
- Only one instance exists.
- Thread-safe initialization using `sync.Once`.

---

# Design Justification

## Factory Pattern

**Purpose:** Centralize ticket object creation.

Why used:

- Decouples object creation from business logic.
- Enables easy addition of new ticket types.
- Follows Open/Closed Principle.

Example extension:

No change required in BookingService.

---

## Builder Pattern

**Purpose:** Construct complex Booking objects step-by-step.

Why used:

- Avoids complex constructors.
- Supports optional fields.
- Improves object consistency.
- Enhances readability.

---

## Singleton Pattern

**Purpose:** Ensure single Logger instance.

Why used:

- Controlled logging access.
- Prevents multiple logger objects.
- Thread-safe and lazy initialization.

---

# SOLID Principles Applied

## Single Responsibility Principle (SRP)

- Ticket → Pricing responsibility
- BookingService → Workflow orchestration
- BookingBuilder → Booking construction
- Repository → Data persistence
- Allocator → Seat management
- Logger → Logging

Each component has one reason to change.

---

## Open / Closed Principle (OCP)

Adding new ticket types:

- Create new Ticket implementation.
- No modification in BookingService.

System is open for extension but closed for modification.

---

## Liskov Substitution Principle (LSP)

All Ticket types:

- Provide valid pricing.
- Can replace one another without breaking behavior.

---

## Interface Segregation Principle (ISP)

Small, focused interfaces:

- PaymentGateway
- NotificationService
- SeatAllocationStrategy
- BookingRepository
- TicketFactory
- BookingBuilder

Avoids unnecessary dependencies.

---

## Dependency Inversion Principle (DIP)

BookingService depends on abstractions:

- TicketFactory
- BookingBuilder
- Strategy interfaces

Concrete implementations are injected via constructor.

This ensures:

- Loose coupling
- Easy testing
- High flexibility

---

# Design Improvement Over Previous Version

Now supports:

- Mixed seat categories
- Flexible pricing
- Real-world booking behavior

---

# Architectural Strengths

- Extensible ticket model
- Non-blocking notification handling
- Constructor-based dependency injection
- Pattern-driven modular design
- Realistic booking workflow

---

# Conclusion

This system demonstrates:

- Effective use of Go interfaces
- Practical implementation of Factory, Builder, and Singleton patterns
- Strong adherence to SOLID principles
- Support for mixed seat type pricing
- Clean, scalable, and maintainable architecture

The design models real-world backend booking systems and is structured for extensibility and clarity.
