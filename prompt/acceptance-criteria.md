# Game Review Application — Engineering Rules

This document defines the engineering rules, architecture guidelines, coding standards, testing requirements, and acceptance criteria for the project.

The AI coding agent MUST follow these rules for all future implementation work.

---

# 1. Project Overview

This is a small fullstack Game Review application.

The project is a monorepo containing two independent applications:

```text
game-review/
├── frontend/
├── backend/
├── docker-compose.yml
└── README.md
```

## Frontend

- Next.js
- TypeScript
- React

## Backend

- Go
- Fiber

## Communication

Frontend and backend communicate exclusively through REST APIs.

## Storage

No external database.

Use in-memory storage or a simple local file.

---

# 2. Core Product Requirements

The application must allow a user to:

1. View a list of games.
2. Select a game.
3. View game details.
4. View all reviews for the selected game.
5. Submit a review.
6. See the newly submitted review immediately.
7. Start with predefined example games and reviews.

The application does NOT require:

- Authentication
- User accounts
- Authorization
- External database
- Payment
- Admin dashboard
- Real-time WebSocket communication
- External game APIs

Do not implement features outside the requirements unless explicitly requested.

---

# 3. Engineering Principles

Follow these principles throughout the project.

## Keep It Simple

This is a small take-home project.

Prefer:

```text
simple + explicit + maintainable
```

over:

```text
complex + abstract + over-engineered
```

Do not introduce abstractions unless they solve an actual problem.

---

## Separation of Concerns

Each layer must have one clear responsibility.

Avoid:

- Business logic inside HTTP handlers.
- API calls directly inside UI components.
- Data access directly inside services.
- Large components containing unrelated responsibilities.
- Duplicated business rules.

---

## Dependency Direction

Backend dependencies should flow in one direction:

```text
HTTP Handler
     ↓
Service
     ↓
Repository
     ↓
Storage
```

Handlers must not directly access repositories.

Services must not depend on Fiber-specific request/response objects.

Repositories must not contain HTTP/business logic.

---

# 4. Backend Architecture

Use a pragmatic layered architecture:

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Storage
```

The backend should be organized around domain concepts.

Recommended structure:

```text
backend/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── game/
│   │   ├── model.go
│   │   ├── dto.go
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── repository_memory.go
│   │
│   ├── review/
│   │   ├── model.go
│   │   ├── dto.go
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── repository_memory.go
│   │
│   └── server/
│       └── router.go
│
├── tests/
│   └── ...
│
├── Dockerfile
├── go.mod
├── go.sum
└── ...
```

The exact structure may be adjusted if there is a strong technical reason.

Do not create unnecessary folders such as:

```text
backend/
├── domain/
├── application/
├── infrastructure/
├── usecase/
├── ports/
├── adapters/
├── providers/
└── factories/
```

unless they become genuinely necessary.

For this project, that would likely be over-engineering.

---

# 5. Backend Layer Responsibilities

## Handler

Responsible for:

- HTTP request parsing
- Input validation at HTTP boundary
- Calling services
- Mapping service results to HTTP responses
- HTTP status codes
- HTTP error responses

Handler MUST NOT:

- Access repository directly
- Implement business rules
- Manipulate storage
- Contain complex domain logic

Example:

```text
POST /api/games/:id/reviews

Handler
  ↓
parse request
  ↓
validate request
  ↓
ReviewService.CreateReview(...)
  ↓
return HTTP response
```

---

# 6. Service Layer

Services contain business logic.

Example responsibilities:

- Verify that a game exists before creating a review.
- Validate rating rules.
- Create review objects.
- Apply business rules.
- Coordinate multiple repositories when necessary.

Services MUST NOT depend on Fiber request/response types.

Services should be testable without starting an HTTP server.

---

# 7. Repository Layer

Repositories are responsible only for data access.

Example:

```go
type GameRepository interface {
    GetAll(ctx context.Context) ([]Game, error)
    GetByID(ctx context.Context, id string) (*Game, error)
}

type ReviewRepository interface {
    GetByGameID(ctx context.Context, gameID string) ([]Review, error)
    Create(ctx context.Context, review Review) (*Review, error)
}
```

The actual implementation can use in-memory storage:

```text
GameRepository
      ↓
MemoryGameRepository

ReviewRepository
      ↓
MemoryReviewRepository
```

Do not introduce a database.

---

# 8. Backend Models

Game and Review must be separate domain models.

Example:

```text
Game
- ID
- Title
- Genre
- Platform

Review
- ID
- GameID
- ReviewerName
- Text
- Rating
- CreatedAt
```

Review must reference the game through `GameID`.

Do not duplicate complete Game objects inside Review.

---

# 9. API Design

The API should follow REST conventions.

Recommended endpoints:

```text
GET  /api/games
GET  /api/games/:id
GET  /api/games/:id/reviews
POST /api/games/:id/reviews
```

Health endpoint:

```text
GET /health
```

Recommended HTTP behavior:

```text
200 OK
201 Created
400 Bad Request
404 Not Found
500 Internal Server Error
```

Do not return `200 OK` for errors simply because the application can technically respond.

---

# 10. API Validation

Review creation must validate:

```text
reviewerName
text
rating
```

At minimum:

- Reviewer name must not be empty.
- Review text must not be empty.
- Rating must be within the chosen rating scale.
- Game must exist.

The validation rules must be tested.

---

# 11. Concurrency

The backend uses in-memory storage.

Because HTTP requests can happen concurrently, shared in-memory data MUST be protected from data races.

Use an appropriate synchronization mechanism such as:

```go
sync.RWMutex
```

where necessary.

Do not ignore concurrent access simply because this is a take-home project.

The code should pass Go's race detector where applicable.

---

# 12. Frontend Architecture

Use Next.js + TypeScript.

Recommended structure:

```text
frontend/
├── app/
│   ├── page.tsx
│   └── games/
│       └── [id]/
│           └── page.tsx
│
├── components/
│   ├── games/
│   │   ├── GameList.tsx
│   │   ├── GameCard.tsx
│   │   └── GameDetails.tsx
│   │
│   └── reviews/
│       ├── ReviewList.tsx
│       ├── ReviewCard.tsx
│       └── ReviewForm.tsx
│
├── lib/
│   └── api/
│       ├── client.ts
│       ├── games.ts
│       └── reviews.ts
│
├── types/
│   ├── game.ts
│   └── review.ts
│
├── hooks/
│   ├── useGames.ts
│   ├── useGame.ts
│   └── useReviews.ts
│
├── tests/
│   └── ...
│
├── public/
├── Dockerfile
├── package.json
└── ...
```

This structure is a guideline, not an absolute requirement.

---

# 13. Frontend Responsibilities

## Pages

Pages should primarily compose application UI and data requirements.

Avoid putting large amounts of business logic directly into:

```text
page.tsx
```

---

## Components

Components should focus on rendering UI and handling user interaction.

Avoid putting raw HTTP requests directly into components.

Bad:

```tsx
const response = await fetch("/api/games");
```

inside a deeply nested UI component.

Prefer:

```text
Component
   ↓
Hook
↓
Service
 ↓
api/routes.ts
   ↓
Backend REST API
```

---

# 14. Frontend API Layer

All backend communication should go through a centralized API client.

Example:

```text
lib/api/client.ts
lib/api/games.ts
lib/api/reviews.ts
```

Do not scatter `fetch()` calls throughout components.

The API layer should handle:

- Base URL
- HTTP methods
- Request headers
- Response parsing
- API errors

---

# 15. TypeScript Rules

Do not use:

```ts
any;
```

unless there is a strong and documented reason.

Prefer explicit types for:

- API responses
- API requests
- Game
- Review
- Error responses
- Component props

Frontend types should reflect backend API contracts.

---

# 16. State Management

Do not introduce Redux, Zustand, or another global state library unless genuinely required.

This application is small enough to use:

- React state
- Server Components where appropriate
- Client Components only when interaction/state is required
- Simple hooks

Keep state close to where it is needed.

---

# 17. Review Submission Behavior

After submitting a review:

```text
User submits form
       ↓
POST /api/games/:id/reviews
       ↓
Backend creates review
       ↓
Frontend receives created review
       ↓
UI updates immediately
```

Do not require:

```text
page reload
application restart
manual refresh
```

to display the newly created review.

Prefer updating the local review state using the created review returned by the API rather than unnecessarily refetching the entire page.

---

# 18. Error Handling

Both frontend and backend must handle expected errors gracefully.

Backend:

```text
Invalid input
Game not found
Unexpected internal error
```

Frontend:

```text
Failed to load games
Failed to load game
Failed to load reviews
Failed to submit review
```

Do not silently swallow errors.

Users should receive a meaningful UI state.

---

# 19. Loading States

The frontend should provide appropriate loading states for asynchronous operations.

At minimum:

- Loading games
- Loading game details/reviews
- Submitting review

Avoid leaving the UI apparently frozen while waiting for the backend.

---

# 20. Testing Requirements

Automated testing is a REQUIRED part of the project.

A feature should not be considered complete without appropriate tests.

---

## Backend Tests

Backend tests are mandatory.

Prioritize testing the service layer.

At minimum test:

### Games

- Get all games.
- Get game by ID.
- Game not found.

### Reviews

- Get reviews for a game.
- Create valid review.
- Reject invalid reviewer name.
- Reject invalid review text.
- Reject invalid rating.
- Reject review for non-existent game.

### Repository

Test important in-memory repository behavior where appropriate.

---

# 21. Backend HTTP Tests

Add API/handler tests for important endpoints.

At minimum:

```text
GET /api/games
GET /api/games/:id
GET /api/games/:id/reviews
POST /api/games/:id/reviews
```

Verify:

- HTTP status
- Response structure
- Validation errors
- Not-found behavior
- Successful creation

---

# 22. Frontend Tests

Frontend tests are strongly expected.

At minimum test important user behavior:

### Game List

- Games are rendered.
- Empty/error state behaves correctly.

### Game Details

- Game information is rendered.
- Reviews are rendered.

### Review Form

- User can enter reviewer name.
- User can enter review text.
- User can select/enter rating.
- Invalid submission is prevented or displays validation.
- Successful submission updates the review list.
- Submission errors are displayed.

Do not test implementation details unnecessarily.

Prefer testing behavior from the user's perspective.

---

# 23. Test Quality

Do not write tests merely to increase coverage.

Bad:

```text
Test that a component exists.
Test that a function was called internally.
```

Prefer:

```text
Given a valid review,
when the user submits it,
then the new review appears in the review list.
```

Tests should describe expected behavior.

---

# 24. Test Commands

The repository must provide simple commands for running tests.

The README must clearly document:

```bash
# Backend
go test ./...

# Frontend
npm test
```

Use the actual commands appropriate to the chosen testing tools.

If possible, backend tests should also support:

```bash
go test -race ./...
```

---

# 25. Seed Data

The application must start with meaningful example data.

Use several games, for example:

```text
Elden Ring
The Witcher 3
Cyberpunk 2077
```

Each game should have at least one example review.

Do not use lorem ipsum.

The seed data should demonstrate the application's functionality.

---

# 26. Docker Rules

Both frontend and backend must have their own Dockerfile.

```text
frontend/Dockerfile
backend/Dockerfile
```

`docker-compose.yml` must orchestrate both services.

Expected developer experience:

```bash
docker compose up --build
```

The application should be usable after this command completes.

---

# 27. Docker Production Awareness

Use multi-stage builds where appropriate.

The final runtime image should contain only what is necessary to run the application.

Do not optimize Docker images excessively at the expense of readability.

The Docker configuration should be easy for another engineer to understand.

---

# 28. Environment Configuration

Do not hardcode environment-specific configuration.

For example:

```text
BACKEND_URL
NEXT_PUBLIC_API_URL
PORT
```

Use environment variables where appropriate.

Do not commit secrets.

There should be no secrets in this project.

---

# 29. Code Style — Go

Follow idiomatic Go.

Prefer:

```text
small functions
explicit error handling
clear naming
simple interfaces
```

Avoid:

```text
giant functions
deep nesting
unnecessary abstractions
global mutable state
magic values
```

Use `gofmt`.

Code should be compatible with standard Go tooling.

---

# 30. Code Style — TypeScript

Follow idiomatic TypeScript/React.

Prefer:

```text
small components
explicit props
reusable UI components
clear naming
simple hooks
```

Avoid:

```text
huge page components
duplicated API calls
any
unnecessary global state
unnecessary abstractions
```

Use the project's formatter/linter consistently.

---

# 31. Naming

Names should describe intent.

Prefer:

```text
CreateReview
GetReviewsByGameID
ReviewForm
ReviewList
getGame
createReview
```

Avoid vague names:

```text
process()
handle()
doStuff()
data()
helper()
utils()
```

unless the context genuinely makes the name meaningful.

---

# 32. No Premature Abstraction

Do not create abstractions for hypothetical future requirements.

For example, do NOT create:

```text
BaseRepository
BaseService
GenericCRUDService
GenericController
AbstractAPIClientFactory
```

unless there is a real requirement for them.

The project is intentionally small.

---

# 33. No Unnecessary Dependencies

Before adding a package/library, ask:

1. Is it required?
2. Does it significantly simplify the code?
3. Could the functionality reasonably be implemented with the existing stack?

Prefer fewer dependencies.

---

# 34. API Contract Consistency

Frontend and backend must agree on:

- Field names
- Data types
- IDs
- Rating scale
- Error format
- HTTP status codes

Do not manually duplicate incompatible assumptions.

When an API contract changes, update both sides and their tests.

---

# 35. Definition of Done

A feature is considered complete only when:

- [ ] Backend implementation is complete.
- [ ] Frontend implementation is complete.
- [ ] Business logic is in the service layer.
- [ ] Data access is in the repository layer.
- [ ] API handler is thin.
- [ ] Input validation exists.
- [ ] Error handling exists.
- [ ] Automated tests exist.
- [ ] Relevant tests pass.
- [ ] No unnecessary `any` exists in TypeScript.
- [ ] No obvious code duplication exists.
- [ ] Code is formatted.
- [ ] Linting passes where configured.
- [ ] Docker build succeeds.
- [ ] Docker Compose starts successfully.
- [ ] README is updated if behavior or commands changed.

---

# 36. Final Acceptance Criteria

The final application must satisfy ALL of the following:

## Functional

- [ ] User can see a list of games.
- [ ] Games have titles.
- [ ] User can select a game.
- [ ] User can see game details.
- [ ] User can see all reviews for a game.
- [ ] Reviews display reviewer name.
- [ ] Reviews display review text.
- [ ] Reviews display rating.
- [ ] User can submit a review.
- [ ] Reviewer name is validated.
- [ ] Review text is validated.
- [ ] Rating is validated.
- [ ] Review can only be created for an existing game.
- [ ] New review appears immediately after successful submission.
- [ ] Application starts with example games.
- [ ] Application starts with example reviews.

## Backend

- [ ] REST API exists.
- [ ] Game and Review are separate models.
- [ ] Handler, Service, and Repository responsibilities are separated.
- [ ] Business logic is not inside handlers.
- [ ] Data access is not inside handlers/services.
- [ ] In-memory/local storage is used.
- [ ] Concurrent in-memory access is safe.
- [ ] Appropriate HTTP status codes are returned.
- [ ] Backend tests pass.

## Frontend

- [ ] Next.js + TypeScript is used.
- [ ] API calls are separated from UI components.
- [ ] Components are reasonably small.
- [ ] Loading states exist.
- [ ] Error states exist.
- [ ] Review submission works.
- [ ] Newly created review appears immediately.
- [ ] Frontend tests cover important user behavior.

## Infrastructure

- [ ] Frontend has a Dockerfile.
- [ ] Backend has a Dockerfile.
- [ ] Docker Compose runs both applications.
- [ ] `docker compose up --build` works from a clean checkout.
- [ ] No external database is required.
- [ ] No secrets are committed.

## Documentation

- [ ] README explains project architecture.
- [ ] README explains prerequisites if any.
- [ ] README explains how to run locally.
- [ ] README explains how to run with Docker.
- [ ] README explains how to run backend tests.
- [ ] README explains how to run frontend tests.
- [ ] README documents important architectural decisions.
- [ ] README mentions known limitations or improvements that would be made with more time.

---

# 37. Agent Behavior

Before modifying code:

1. Inspect the existing repository.
2. Understand the current architecture.
3. Reuse existing patterns where appropriate.
4. Do not overwrite working code unnecessarily.
5. Check whether the requested functionality already exists.
6. Make the smallest reasonable change.

After modifying code:

1. Format the code.
2. Run relevant tests.
3. Run lint/type checks where available.
4. Verify Docker configuration if infrastructure was changed.
5. Report what changed.
6. Report tests that were run and their result.
7. Clearly mention anything that could not be verified.

Never claim that a test or command passed if it was not actually executed.

---

# 38. Priority Order

When making engineering decisions, prioritize:

```text
1. Correctness
2. Acceptance criteria
3. Maintainability
4. Testability
5. Simplicity
6. Developer experience
7. Performance
8. Visual polish
```

Do not sacrifice correctness or maintainability for visual polish.

The goal is to demonstrate strong engineering judgment through a small, clean, understandable application.
