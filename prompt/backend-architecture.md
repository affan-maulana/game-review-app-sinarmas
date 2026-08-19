# Backend Architecture

This document defines the architecture and engineering conventions for the Game Review backend.

The backend is built with:

- Go
- Fiber
- REST API
- In-memory storage

The architecture follows a **modular design**.

Each business domain is isolated into its own module, while shared application infrastructure is kept under `internal`.

---

# 1. Architecture Overview

The backend follows this high-level structure:

```text
HTTP Request
     ↓
   Routes
     ↓
  Handler
     ↓
  Service
     ↓
 Repository
     ↓
 In-Memory Storage
```

The application itself is organized into modules:

```text
internal/
└── modules/
    ├── game/
    └── review/
```

Each module owns its:

- Models
- DTOs
- Repository
- Service
- Handler

Shared infrastructure is kept outside the modules.

---

# 2. Recommended Folder Structure

Baik, ini justru **lebih clean dan lebih realistis untuk take-home project** karena:

- tidak terlalu banyak nested folder
- lebih cepat navigasi saat coding
- lebih mudah dijelaskan saat interview
- mirip struktur production kecil–menengah di Go

Saya akan revisi bagian `modules/` sesuai preferensi kamu: **semua layer (dto, handler, model, service, repository) berada langsung di dalam folder module**.

---

### ✅ Revisi struktur `modules/`

```text
internal/
└── modules/
    ├── game/
    │   ├── game_dto.go
    │   ├── game_handler.go
    │   ├── game_model.go
    │   ├── game_repository.go
    │   ├── game_repository_memory.go
    │   └── game_service.go
    │
    └── review/
        ├── review_dto.go
        ├── review_handler.go
        ├── review_model.go
        ├── review_repository.go
        ├── review_repository_memory.go
        ├── review_service.go
        └── review_service_impl.go
```

---

### 🧠 Penyesuaian konsep (penting untuk interview)

Dengan struktur ini, kamu bisa jelaskan seperti ini:

> “Each module is self-contained. Instead of splitting by layer folders, I group everything by domain to improve locality and reduce navigation overhead.”

---

### 📌 Penjelasan tiap file (biar kamu siap interview)

#### `model.go`

Domain entity (pure business object)

#### `dto.go`

Request/response contract API

#### `handler.go`

HTTP layer (Fiber)

#### `service.go`

Interface service (contract)

#### `game_service.go / review_service.go`

Business logic implementation

#### `repository.go`

Interface repository

#### `repository_memory.go`

In-memory implementation

---

### ⚖️ Kenapa ini lebih bagus untuk case kamu

✅ Lebih simple (tidak over-engineered)
✅ Lebih cepat develop (important untuk take-home)
✅ Lebih mudah di-review interviewer
✅ Tetap clean architecture secara konsep
✅ Cocok untuk project tanpa database

---

### 🚀 Bonus insight (ini bisa jadi nilai plus di interview)

Kalau ditanya:

> “Kenapa tidak dipisah folder dto/handler/service?”

Jawaban ideal:

> “Because the project is small and domain-focused, I prefer vertical slicing per module. It improves cohesion and reduces context switching. If the system grows, I can always split layers into separate packages.”

---

Kalau kamu mau, next step aku bisa bantu:

✅ refactor `ARCHITECTURE.md` sesuai struktur ini
✅ atau bikin **contoh real flow Game + Review (end-to-end code design)**
✅ atau bantu bikin **jawaban interview untuk architecture decision ini**

Tinggal bilang 👍

The exact filenames may change as the application evolves.

The important principle is that **business functionality is organized by module**.

---

# 3. `cmd/`

The `cmd` directory contains application entry points.

```text
cmd/
└── server/
    └── main.go
```

`main.go` should remain thin.

Its responsibility is to:

1. Load configuration.
2. Initialize dependencies/container.
3. Initialize the Fiber application.
4. Register middleware.
5. Register routes.
6. Start the HTTP server.

Avoid putting business logic inside `main.go`.

Conceptually:

```text
main.go
   ↓
config
   ↓
container
   ↓
fiber app
   ↓
middleware
   ↓
routes
   ↓
server
```

---

# 4. `internal/`

The `internal` directory contains application-specific code that should not be imported by external applications.

It contains:

```text
internal/
├── config/
├── container/
├── middleware/
├── modules/
└── routes/
```

These packages belong specifically to this application.

---

# 5. `internal/config/`

Configuration and infrastructure setup live here.

Example:

```text
config/
├── config.go
└── db.go
```

Responsibilities may include:

- Environment configuration.
- Application configuration.
- Storage initialization.
- Database initialization if a database existed.

For this project there is no external database.

Therefore `db.go` should not introduce PostgreSQL/MySQL/etc.

If needed, it can contain initialization for the in-memory data store.

Example conceptual flow:

```text
config.Load()
      ↓
storage initialization
      ↓
container
```

Keep environment-specific configuration out of business modules.

---

# 6. `internal/container/`

The container is responsible for wiring application dependencies.

Example:

```text
Container
├── GameRepository
├── GameService
├── GameHandler
├── ReviewRepository
├── ReviewService
└── ReviewHandler
```

Conceptually:

```text
Repository
     ↓
  Service
     ↓
  Handler
```

The container creates these dependencies and passes them to the appropriate components.

The container should NOT:

- Contain business logic.
- Perform HTTP requests.
- Implement repository logic.
- Become a generic dependency injection framework.

Keep it simple.

---

# 7. `internal/middleware/`

Contains HTTP middleware shared across the application.

Examples:

```text
middleware/
├── logger.go
├── recover.go
└── cors.go
```

Possible middleware:

- Request logging.
- Panic recovery.
- CORS.
- Request ID.

Only add middleware that has an actual purpose.

Do not add authentication middleware because authentication is not part of the requirements.

---

# 8. `internal/modules/`

This is the main business layer of the application.

Each business domain gets its own module.

Current modules:

```text
modules/
├── game/
└── review/
```

A module should be as self-contained as reasonably possible.

---

# 9. Module Structure

Each module follows:

```text
module/
├── dto/
├── handler/
├── model/
├── repository/
└── service/
```

Dependency direction:

```text
Handler
   ↓
Service
   ↓
Repository
```

The model represents the domain data.

DTOs represent HTTP/API input and output structures.

---

# 10. Model

Location:

```text
modules/game/model/
modules/review/model/
```

Models represent domain entities.

Example Game:

```go
type Game struct {
    ID       string
    Title    string
    Genre    string
    Platform string
}
```

Example Review:

```go
type Review struct {
    ID           string
    GameID       string
    ReviewerName string
    Text         string
    Rating       int
    CreatedAt    time.Time
}
```

Game and Review must remain separate domain concepts.

Review references a Game using `GameID`.

---

# 11. DTO

Location:

```text
modules/game/dto/
modules/review/dto/
```

DTOs represent data crossing the API boundary.

For example:

```go
type CreateReviewRequest struct {
    ReviewerName string `json:"reviewerName"`
    Text         string `json:"text"`
    Rating       int    `json:"rating"`
}
```

DTOs should not automatically be treated as domain models.

This separation allows API contracts to evolve independently from internal domain structures.

---

# 12. Handler Layer

Location:

```text
modules/game/handler/
modules/review/handler/
```

Handlers are responsible for HTTP concerns.

Responsibilities:

- Read route parameters.
- Parse request body.
- Validate request format.
- Call service methods.
- Map errors to HTTP responses.
- Return HTTP responses.

Handlers should remain thin.

Example:

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Response
```

Handlers MUST NOT:

- Access repositories directly.
- Implement business rules.
- Manipulate storage.
- Contain complex domain logic.

Bad:

```text
Handler
 ├── find game
 ├── validate rating
 ├── create review
 ├── modify storage
 └── return response
```

Preferred:

```text
Handler
   ↓
ReviewService.Create(...)
```

---

# 13. Service Layer

Location:

```text
modules/game/service/
modules/review/service/
```

Services contain business logic.

Examples:

### Game Service

```text
GetGames
GetGameByID
```

### Review Service

```text
GetReviewsByGameID
CreateReview
```

The Review service may need to verify that the referenced game exists before creating a review.

Conceptually:

```text
CreateReview
     ↓
validate review
     ↓
check game exists
     ↓
create review
     ↓
repository.Create()
```

Services should not depend on Fiber-specific types.

Avoid:

```go
func (s *Service) Create(c *fiber.Ctx)
```

Prefer plain application-level parameters and return values.

---

# 14. Repository Layer

Location:

```text
modules/game/repository/
modules/review/repository/
```

Repositories are responsible for data access.

For this project, repositories use in-memory storage.

Example:

```text
GameRepository
      ↓
MemoryGameRepository

ReviewRepository
      ↓
MemoryReviewRepository
```

The repository interface defines what the service needs.

Example:

```go
type Repository interface {
    GetAll(ctx context.Context) ([]model.Game, error)
    GetByID(ctx context.Context, id string) (*model.Game, error)
}
```

The service should depend on the interface rather than the concrete implementation.

This makes the service easy to test.

---

# 15. In-Memory Storage

Because the exercise explicitly prohibits an external database, use in-memory storage.

Example:

```text
MemoryGameRepository
    └── []Game

MemoryReviewRepository
    └── []Review
```

Because HTTP requests may execute concurrently, shared mutable state must be protected.

Use synchronization such as:

```go
sync.RWMutex
```

where necessary.

The implementation should be safe against concurrent reads/writes.

---

# 16. Routes

Location:

```text
internal/routes/
```

Routes are separated by module.

Example:

```text
routes/
├── routes.go
├── game_routes.go
└── review_routes.go
```

Game routes:

```text
GET /api/games
GET /api/games/:id
```

Review routes:

```text
GET  /api/games/:id/reviews
POST /api/games/:id/reviews
```

The routes package is responsible for connecting HTTP endpoints to handlers.

It should not contain business logic.

---

# 17. Route Registration

Conceptually:

```text
main.go
   ↓
routes.Register(...)
   ↓
game routes
review routes
```

Example:

```text
/api
 ├── /games
 │    ├── GET /
 │    └── GET /:id
 │
 └── /games/:id/reviews
      ├── GET /
      └── POST /
```

Keep route definitions readable.

---

# 18. `pkg/`

The `pkg` directory contains reusable packages that are not specific to a particular business module.

Examples:

```text
pkg/
├── response/
│   └── response.go
└── validator/
    └── validator.go
```

A package belongs in `pkg` only when it represents genuinely reusable application infrastructure.

Do not use `pkg` as a dumping ground for miscellaneous business logic.

---

# 19. Response Package

A shared response helper can live under:

```text
pkg/response/
```

Example:

```go
func Success(c *fiber.Ctx, statusCode int, data interface{}) error {
    return c.Status(statusCode).JSON(APIResponse{
        Status: "success",
        Data:   data,
    })
}
```

Error responses should follow a consistent structure as well.

For example:

```json
{
  "status": "error",
  "message": "Game not found"
}
```

The response package should only deal with HTTP response formatting.

It should not contain business logic.

---

# 20. Validator Package

Reusable validation helpers can live under:

```text
pkg/validator/
```

Example:

```go
func TrimString(s *string) *string {
    trimmed := strings.TrimSpace(*s)
    return &trimmed
}
```

Only put genuinely reusable utilities here.

Business-specific validation belongs in the appropriate module/service.

For example:

```text
Rating must be between 1 and 5
```

is a Review business rule and should NOT become:

```text
pkg/validator.ValidateRating()
```

unless there is a genuine cross-domain reason.

---

# 21. Migration Directory

The project does not use an external database.

Therefore:

```text
migration/
```

may remain empty or contain future migration placeholders.

Do not introduce a database just to make the migration directory useful.

If there is no migration requirement, it is acceptable for this directory not to contain active migration files.

---

# 22. Error Handling

Errors should be explicit and predictable.

The service layer should return meaningful errors.

For example:

```text
GameNotFound
ReviewNotFound
InvalidReview
```

The handler maps these errors into HTTP responses.

Example:

```text
Service
   ↓
ErrGameNotFound
   ↓
Handler
   ↓
404 Not Found
```

Do not put HTTP status codes inside domain/business logic unless there is a strong reason.

---

# 23. HTTP Status Codes

Use appropriate status codes.

Recommended:

```text
GET successful       → 200 OK
POST successful      → 201 Created
Invalid request      → 400 Bad Request
Resource not found   → 404 Not Found
Unexpected error     → 500 Internal Server Error
```

Do not return `200 OK` for failed operations.

---

# 24. Validation

Review creation must validate:

```text
ReviewerName
Text
Rating
GameID
```

At minimum:

```text
Reviewer name must not be empty.
Review text must not be empty.
Rating must be within the selected scale.
Game must exist.
```

Validation can occur at multiple boundaries:

```text
Handler
  ↓
request shape / parsing validation

Service
  ↓
business validation

Repository
  ↓
data access
```

The service layer remains responsible for business rules.

---

# 25. Testing Strategy

Automated testing is REQUIRED.

Tests should focus primarily on business behavior.

Recommended priority:

```text
1. Service tests
2. Handler/API tests
3. Repository tests
```

The exact test structure can follow the module structure.

For example:

```text
modules/
└── review/
    ├── handler/
    │   └── review_handler_test.go
    ├── service/
    │   └── review_service_test.go
    └── repository/
        └── memory_test.go
```

---

# 26. Service Tests

Service tests are the most important backend tests.

Review service should test:

- Creating a valid review.
- Rejecting empty reviewer name.
- Rejecting empty review text.
- Rejecting invalid rating.
- Rejecting a review for a non-existent game.
- Successfully retrieving reviews.

Game service should test:

- Retrieving all games.
- Retrieving an existing game.
- Handling a non-existent game.

Services should be testable without starting Fiber.

---

# 27. Repository Tests

Repository tests should verify important storage behavior.

For example:

```text
Create review
Get review
Get reviews by game ID
Get game by ID
Get all games
```

Do not spend excessive effort testing trivial implementation details.

---

# 28. Handler/API Tests

Important HTTP endpoints should have tests.

At minimum:

```text
GET  /api/games
GET  /api/games/:id
GET  /api/games/:id/reviews
POST /api/games/:id/reviews
```

Verify:

- Status codes.
- Response body.
- Validation behavior.
- Not-found behavior.
- Successful creation.

Fiber's testing utilities can be used to test handlers without running a real external server.

---

# 29. Race Detection

Because the application uses in-memory mutable storage, concurrency should be considered.

Where possible, verify with:

```bash
go test -race ./...
```

The application should not introduce obvious data races.

---

# 30. Code Quality

Follow idiomatic Go.

Required:

- Run `gofmt`.
- Keep functions small.
- Handle errors explicitly.
- Use meaningful names.
- Avoid unnecessary interfaces.
- Avoid unnecessary abstractions.
- Avoid global mutable state.
- Avoid magic values.
- Keep dependencies explicit.

---

# 31. Interfaces

Interfaces should be defined where they are consumed.

For example, if the service needs a repository:

```go
type GameRepository interface {
    GetAll(ctx context.Context) ([]model.Game, error)
    GetByID(ctx context.Context, id string) (*model.Game, error)
}
```

The service can depend on this interface.

Avoid creating interfaces for every struct without a testing or architectural reason.

Do not create generic interfaces such as:

```text
BaseRepository
BaseService
CRUDRepository
```

just for abstraction.

---

# 32. Dependency Flow

The dependency direction should remain:

```text
Routes
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
Storage
```

Infrastructure dependencies should not leak into business logic.

For example, the service should not know about:

```text
fiber.Ctx
HTTP headers
HTTP status codes
JSON response formatting
```

---

# 33. Module Independence

Modules should not directly access another module's internal implementation.

For example:

```text
review/service
```

should not access:

```text
game/repository/memory.go
```

Instead, it should depend on an appropriate interface or service contract.

For this project, the Review service needs to verify that a Game exists.

Prefer dependency injection:

```text
ReviewService
    ↓
GameRepository interface
```

rather than directly accessing the Game module's concrete repository.

---

# 34. Container Dependency Wiring

The container is responsible for assembling dependencies.

Conceptually:

```text
GameRepository
       ↓
GameService
       ↓
GameHandler


GameRepository ─────┐
                    ↓
ReviewService
       ↓
ReviewHandler
```

This keeps dependency creation out of handlers and services.

---

# 35. Seed Data

The application must start with example games and reviews.

Example:

```text
Games:
- Elden Ring
- The Witcher 3
- Cyberpunk 2077
```

Each game should have at least one review.

Seed data should be initialized during application startup.

Keep seed data separate from business logic where practical.

---

# 36. API Contract

Recommended endpoints:

```text
GET  /health

GET  /api/games
GET  /api/games/:id
GET  /api/games/:id/reviews
POST /api/games/:id/reviews
```

Create review request:

```json
{
  "reviewerName": "John",
  "text": "Great game!",
  "rating": 5
}
```

Created response should contain the created review.

This allows the frontend to immediately add the new review to its local state.

---

# 37. Docker

The backend must have its own Dockerfile:

```text
backend/Dockerfile
```

Prefer a multi-stage build:

```text
Go source
   ↓
Build stage
   ↓
Compiled binary
   ↓
Small runtime image
```

The backend should be runnable through Docker Compose.

---

# 38. Definition of Done

A backend feature is complete only when:

- [ ] Handler is implemented.
- [ ] Service contains business logic.
- [ ] Repository handles data access.
- [ ] Dependencies are wired through the container.
- [ ] Routes are registered.
- [ ] Input validation exists.
- [ ] Appropriate errors are returned.
- [ ] Appropriate HTTP status codes are returned.
- [ ] Automated tests exist.
- [ ] Tests pass.
- [ ] `go test ./...` passes.
- [ ] `go test -race ./...` passes where applicable.
- [ ] `gofmt` has been applied.
- [ ] Docker build succeeds.
- [ ] README is updated when necessary.

---

# 39. Backend Acceptance Criteria

## Game

- [ ] `GET /api/games` returns all seeded games.
- [ ] `GET /api/games/:id` returns a game.
- [ ] Non-existent game returns `404`.
- [ ] Game contains at least an ID and title.

## Reviews

- [ ] `GET /api/games/:id/reviews` returns reviews for the selected game.
- [ ] Non-existent game returns `404`.
- [ ] `POST /api/games/:id/reviews` creates a review.
- [ ] Created review contains reviewer name.
- [ ] Created review contains review text.
- [ ] Created review contains rating.
- [ ] Created review contains a game ID.
- [ ] Created review contains a creation timestamp.
- [ ] Invalid reviewer name is rejected.
- [ ] Invalid review text is rejected.
- [ ] Invalid rating is rejected.
- [ ] Review for a non-existent game is rejected.
- [ ] Created review is available immediately through subsequent requests.

## Architecture

- [ ] Routes are separated by module.
- [ ] Handlers do not access repositories directly.
- [ ] Business logic lives in services.
- [ ] Data access lives in repositories.
- [ ] Modules remain reasonably isolated.
- [ ] Shared HTTP utilities live in `pkg`.
- [ ] Dependencies are wired through the container.
- [ ] No unnecessary external infrastructure is introduced.

## Testing

- [ ] Service tests exist.
- [ ] Repository tests exist where useful.
- [ ] HTTP/API tests exist for important endpoints.
- [ ] Validation behavior is tested.
- [ ] Error behavior is tested.
- [ ] Tests pass.
- [ ] No obvious data races exist.

---

# 40. Architecture Summary

The final backend architecture should be understandable at a glance:

```text
                    ┌─────────────────┐
                    │      Routes     │
                    │  Game / Review  │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │     Handler     │
                    │  HTTP concerns  │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │     Service     │
                    │ Business Logic  │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   Repository    │
                    │   Data Access   │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ In-Memory Store │
                    └─────────────────┘
```

At the application level:

```text
cmd/
  ↓
container/
  ↓
routes/
  ↓
modules/
  ├── game
  └── review

Shared infrastructure:
  ├── config
  ├── middleware
  └── pkg
```

The architecture intentionally uses **modular design + handler/service/repository separation** without introducing unnecessary enterprise patterns.

The goal is to demonstrate that the application can remain small while still being:

- Maintainable
- Testable
- Modular
- Easy to understand
- Easy to extend
- Easy to explain during the technical interview
