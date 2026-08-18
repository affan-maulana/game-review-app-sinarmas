# Game Review Application

A fullstack Game Review application built with **Next.js + TypeScript** frontend and **Go + Fiber** backend.

## Architecture

```
game-review-app/
├── frontend/          # Next.js + TypeScript
│   ├── app/           # Pages (App Router)
│   ├── components/    # UI components (games/, reviews/)
│   ├── hooks/         # Custom React hooks
│   ├── services/      # Application services
│   ├── lib/api/       # API client and routes
│   └── types/         # TypeScript type definitions
│
├── backend/           # Go + Fiber
│   ├── cmd/server/    # Entry point
│   ├── internal/
│   │   ├── game/      # Game module (model, handler, service, repository)
│   │   ├── review/    # Review module (model, handler, service, repository)
│   │   ├── config/    # Configuration and seed data
│   │   ├── container/ # Dependency wiring
│   │   ├── middleware/ # HTTP middleware (CORS, logger, recover)
│   │   └── server/    # Route registration
│   └── pkg/           # Shared utilities (response, validator)
│
├── docker-compose.yml
└── README.md
```

### Backend Layer Flow

```
HTTP Request → Routes → Handler → Service → Repository → In-Memory Storage
```

- **Handlers**: HTTP parsing, input validation, response mapping
- **Services**: Business logic, validation rules
- **Repositories**: Data access (in-memory with sync.RWMutex)

### Frontend Layer Flow

```
UI Components → Hooks → Services → API Client → Go Fiber REST API
```

- **Components**: UI rendering, user interaction
- **Hooks**: State management, loading/error states
- **Services**: Application operations
- **API Client**: HTTP communication

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/games` | List all games |
| GET | `/api/games/:id` | Get game details |
| GET | `/api/games/:id/reviews` | Get reviews for a game |
| POST | `/api/games/:id/reviews` | Create a review |

### Create Review Request

```json
{
  "reviewerName": "John",
  "text": "Great game!",
  "rating": 5
}
```

## Prerequisites

- **Go 1.21+**
- **Node.js 18+**
- **Docker & Docker Compose** (optional)

## Running Locally

### Backend

```bash
cd backend
go mod tidy
go run ./cmd/server
```

Backend runs on `http://localhost:3001`

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on `http://localhost:3000`

## Running with Docker

```bash
docker compose up --build
```

- Frontend: `http://localhost:3000`
- Backend: `http://localhost:3001`

## Running Tests

### Backend Tests

```bash
cd backend
go test ./...
go test -race ./...
```

### Frontend Tests

```bash
cd frontend
npm test
```

## Seed Data

The application starts with 5 example games:

- Elden Ring
- The Witcher 3: Wild Hunt
- Cyberpunk 2077
- God of War Ragnarök
- Baldur's Gate 3

Each game has at least one example review.

## Key Design Decisions

1. **In-Memory Storage**: No external database. Data persists during server runtime only. Protected with `sync.RWMutex` for concurrent access safety.

2. **Module-Based Architecture**: Game and Review are separate domain modules, each with their own model, DTO, handler, service, and repository.

3. **Dependency Injection via Container**: The container wires dependencies (repository → service → handler) and keeps business logic independent of HTTP framework.

4. **Frontend Layer Separation**: Components don't call `fetch()` directly. All API communication goes through the service → API client chain.

5. **No Global State Management**: Uses React hooks and local state. No Redux/Zustand needed for this application size.

## Environment Variables

### Backend

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3001` | Server port |
| `FRONTEND_URL` | `http://localhost:3000` | CORS allowed origin |

### Frontend

| Variable | Default | Description |
|----------|---------|-------------|
| `NEXT_PUBLIC_API_URL` | `http://localhost:3001` | Backend API URL |
