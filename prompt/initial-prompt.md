# Game Review Application — Initial Project Setup

You are working on a take-home fullstack coding exercise.

The repository is a monorepo containing two completely separate applications:

- `frontend/` — Next.js + TypeScript
- `backend/` — Go + Fiber

The frontend and backend must communicate exclusively through a REST API.

Your task in this phase is to **initialize the project architecture and development environment only**. Do not implement the complete application features yet.

---

## Project Goal

Build a small Game Review web application.

A video game reviewer should be able to:

1. See a list of games.
2. Select a game and see its details.
3. Read all reviews for a game.
4. Submit a new review for a game.
5. See the newly submitted review immediately without restarting the application.

The application does not require authentication or user accounts.

There is no external database. Data can initially be stored in memory or a simple local file.

---

## Acceptance Criteria

The final application must satisfy:

### Games

- Show a list of games.
- Every game must have at least a title.
- Additional information such as genre or platform may be included.

### Game Details

- Clicking/selecting a game shows its details.
- The game detail page must show all reviews belonging to that game.

### Reviews

Every review must contain:

- Reviewer name
- Review text
- Rating

Choose a simple rating scale, preferably 1–5.

### Creating Reviews

A user must be able to submit a new review for a game through the frontend.

After successfully submitting a review:

- The new review must appear immediately in the UI.
- The application must not require a restart.
- The backend must persist the review in its in-memory/local storage.

### Initial Data

The application must contain several example games and reviews when it starts so the UI is not empty.

---

# Non-Functional Requirements

## 1. Separate Frontend and Backend

The frontend and backend must remain independent applications.

Expected structure:

```text
/
├── frontend/
├── backend/
├── docker-compose.yml
└── README.md
```

The frontend must communicate with the backend using REST APIs.

Do not put backend logic inside Next.js API routes.

---

## 2. Backend Architecture

Use Go + Fiber.

Keep responsibilities separated.

At minimum, separate:

```text
HTTP handlers/controllers
        ↓
business/service layer
        ↓
data/repository layer
        ↓
in-memory/local storage
```

Game and Review should be separate domain concepts/models.

Avoid putting business logic directly inside HTTP handlers.

Do not introduce unnecessary enterprise architecture or abstractions.

The architecture should be simple enough for a small take-home project but clean enough to demonstrate maintainability.

---

## 3. Frontend Architecture

Use:

- Next.js
- TypeScript

Keep UI components, API communication, and application logic reasonably separated.

Avoid putting all logic inside page components.

Create a small API client/service layer for communicating with the backend.

Use a simple and maintainable component structure.

Do not over-engineer the frontend.

---

## 4. REST API

The backend should expose an API similar to:

```text
GET  /api/games
GET  /api/games/:id
GET  /api/games/:id/reviews
POST /api/games/:id/reviews
```

You may adjust the exact API structure if you have a strong reason, but keep it RESTful and simple.

For creating a review, the request should contain something equivalent to:

```json
{
  "reviewerName": "John",
  "text": "Great game!",
  "rating": 5
}
```

The API should validate input and return appropriate HTTP status codes.

---

## 5. Data Storage

Do NOT add:

- PostgreSQL
- MySQL
- MongoDB
- Redis
- External database services
- ORM

Use in-memory storage or a simple local file.

The purpose of this exercise is not database design.

Seed the application with several games and reviews.

---

## 6. Docker

Both applications must be containerized.

Expected structure:

```text
frontend/Dockerfile
backend/Dockerfile
docker-compose.yml
```

The entire application should be runnable with a single command.

The target developer experience should be:

```bash
docker compose up --build
```

The exact ports can be chosen appropriately.

Make sure the frontend can communicate with the backend correctly when running through Docker Compose.

---

# Testing

Automated tests are important for this exercise.

Set up the project so that tests can be added and run easily.

Backend tests are required.

Frontend tests are strongly preferred.

At minimum, establish the testing infrastructure during this setup phase.

Do not focus on achieving high test coverage yet.

Focus on making the test setup clean and maintainable.

---

# Code Quality Expectations

The reviewer will inspect the source code and discuss architectural decisions during a follow-up call.

Therefore:

- Prefer simple and explicit code.
- Follow idiomatic Go.
- Use TypeScript properly.
- Avoid unnecessary abstractions.
- Avoid premature optimization.
- Keep responsibilities separated.
- Use meaningful names.
- Handle errors explicitly.
- Validate API input.
- Keep configuration/environment variables clear.
- Make the project easy for another engineer to understand.

Do not implement unnecessary features just to make the project look larger.

---

# Initial Setup Scope

For this task, ONLY do the project initialization.

Please:

1. Inspect the repository.
2. Create the `frontend/` Next.js application.
3. Create the `backend/` Go Fiber application.
4. Establish the initial backend architecture.
5. Establish the initial frontend architecture.
6. Add Dockerfiles for both applications.
7. Add `docker-compose.yml`.
8. Configure frontend → backend communication infrastructure.
9. Configure environment variables appropriately.
10. Configure backend testing.
11. Configure frontend testing.
12. Add basic health-check endpoints where appropriate.
13. Add a basic README explaining how to start the development environment.

Do NOT implement the complete game/review feature yet.

---

# Important Constraints

Do not:

- Add authentication.
- Add user accounts.
- Add an external database.
- Add WebSockets unless they become clearly necessary later.
- Add Redis.
- Add Kubernetes.
- Add microservices.
- Add unnecessary libraries.
- Build an overly complex architecture.
- Implement features that are not required by the acceptance criteria.

The goal is a small, professional take-home project.

---

# Expected Result

After the setup is complete, the repository should roughly look like:

```text
game-review/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── game/
│   │   ├── review/
│   │   ├── handler/
│   │   ├── service/
│   │   └── repository/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── ...
│
├── frontend/
│   ├── app/
│   ├── components/
│   ├── lib/
│   ├── ...
│   ├── Dockerfile
│   ├── package.json
│   └── ...
│
├── docker-compose.yml
├── README.md
└── .gitignore
```

The exact directory structure can differ if you have a better simple structure, but responsibilities must remain clearly separated.

Before making implementation decisions, inspect the repository and existing configuration first.

After completing the setup, provide a concise summary of:

- What was created
- Architecture decisions
- Commands to run locally
- Commands to run with Docker
- Commands to run tests
- Any assumptions or decisions that should be discussed before implementing the actual features

Do not proceed to implement the full Game Review functionality unless explicitly asked in the next step.
