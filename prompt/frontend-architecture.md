# Frontend Architecture

This document describes the frontend architecture and engineering conventions for the Game Review application.

The frontend uses **Next.js + TypeScript** and communicates with the Go Fiber backend through a REST API.

The architecture follows this responsibility flow:

```text
UI Components
      ↓
    Hooks
      ↓
   Services
      ↓
  API Client
      ↓
 Go Fiber REST API
```

The goal is to keep UI components focused on presentation, hooks responsible for UI state and orchestration, services responsible for application operations, and the API client responsible for HTTP communication.

---

# 1. Architecture Principles

The frontend should follow these principles:

- Keep components focused on UI and user interaction.
- Keep state management close to where it is needed.
- Keep API communication outside UI components.
- Keep application operations inside services.
- Centralize HTTP communication in the API client.
- Prefer simple and explicit code.
- Avoid unnecessary global state.
- Avoid premature abstractions.
- Keep TypeScript types explicit.
- Make important behavior easy to test.

The architecture should remain lightweight because this is a small application.

---

# 2. Recommended Folder Structure

```text
frontend/
├── app/
│   ├── page.tsx
│   │
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
├── hooks/
│   ├── useGames.ts
│   ├── useGame.ts
│   └── useReviews.ts
│
├── services/
│   ├── game.service.ts
│   └── review.service.ts
│
├── api/
│   ├── client.ts
│   └── routes.ts
│
├── types/
│   ├── game.ts
│   ├── review.ts
│   └── api.ts
│
├── tests/
│   ├── components/
│   ├── hooks/
│   └── services/
│
├── public/
│
├── Dockerfile
├── package.json
├── tsconfig.json
└── ...
```

The structure is a guideline.

Do not create additional layers or folders unless there is a clear reason.

---

# 3. Layer Responsibilities

## UI Layer

Location:

```text
components/
app/
```

Responsibilities:

- Render UI.
- Receive props.
- Handle user interaction.
- Display loading, error, empty, and success states.
- Trigger hooks/actions in response to user interaction.

UI components should NOT:

- Call `fetch()` directly.
- Know the backend URL.
- Construct API URLs.
- Contain business rules.
- Contain large amounts of data-fetching logic.

Example:

```tsx
function ReviewForm() {
  const { createReview, isSubmitting } = useReviews();

  // UI and interaction logic
}
```

The component should not know how the HTTP request is implemented.

---

# 4. Hooks Layer

Location:

```text
hooks/
```

Hooks are responsible for coordinating UI state and application operations.

Examples:

```text
useGames()
useGame()
useReviews()
```

Responsibilities:

- Manage loading state.
- Manage error state.
- Manage local UI/application state.
- Call services.
- Expose actions and state to components.
- Coordinate updates after mutations.

Example:

```text
ReviewForm
     ↓
useReviews()
     ↓
reviewService.createReview()
```

Hooks should NOT:

- Contain raw `fetch()` calls.
- Know HTTP implementation details.
- Directly manipulate API URLs.
- Duplicate service logic.

---

# 5. Services Layer

Location:

```text
services/
```

Services represent application-level operations.

Examples:

```text
game.service.ts
review.service.ts
```

Example operations:

```text
gameService.getGames()
gameService.getGame(id)

reviewService.getReviews(gameId)
reviewService.createReview(gameId, payload)
```

Responsibilities:

- Define application operations.
- Call the API client.
- Transform API data when necessary.
- Provide a clean interface for hooks.

Services should not:

- Render UI.
- Manage React state.
- Access React hooks.
- Know about components.
- Directly manipulate browser UI.

The service should be usable independently from React.

---

# 6. API Client Layer

Location:

```text
api/
```

The API layer is responsible for HTTP communication with the backend.

Recommended structure:

```text
api/
├── client.ts
└── routes.ts
```

## `client.ts`

Responsible for generic HTTP behavior:

- Base URL.
- HTTP methods.
- Headers.
- JSON serialization.
- Response parsing.
- HTTP error handling.

For example:

```text
apiClient.get(...)
apiClient.post(...)
```

The API client should not contain game/review business logic.

---

# 7. API Routes

`routes.ts` contains centralized backend endpoint definitions.

Example:

```ts
export const routes = {
  games: {
    list: "/api/games",
    detail: (id: string) => `/api/games/${id}`,
    reviews: (id: string) => `/api/games/${id}/reviews`,
  },
};
```

The purpose is to avoid scattering API paths throughout the application.

Avoid:

```ts
fetch(`${API_URL}/api/games/${id}/reviews`);
```

inside components or hooks.

Prefer:

```text
Component
    ↓
Hook
    ↓
Service
    ↓
API Client
    ↓
routes.ts
```

---

# 8. Important Terminology

This project uses:

```text
API Client
```

to refer to the frontend HTTP communication layer.

Do NOT create Next.js API Routes such as:

```text
app/api/games/route.ts
```

unless there is a specific requirement to use Next.js as a backend-for-frontend.

The actual backend is Go Fiber.

The intended architecture is:

```text
Browser
   ↓
Next.js Frontend
   ↓
REST API
   ↓
Go Fiber Backend
```

---

# 9. Types

Location:

```text
types/
```

Types should represent the API/domain data used by the frontend.

Example:

```ts
export interface Game {
  id: string;
  title: string;
  genre?: string;
  platform?: string;
}

export interface Review {
  id: string;
  gameId: string;
  reviewerName: string;
  text: string;
  rating: number;
  createdAt: string;
}
```

Avoid duplicating the same type definitions across components, hooks, and services.

---

# 10. Data Flow — Reading Games

When the game list is displayed:

```text
Games Page
    ↓
useGames()
    ↓
gameService.getGames()
    ↓
apiClient.get(routes.games.list)
    ↓
GET /api/games
    ↓
Go Fiber
```

The response flows back in the opposite direction:

```text
Go Fiber
    ↓
API Client
    ↓
Service
    ↓
Hook
    ↓
UI
```

---

# 11. Data Flow — Game Details

```text
Game Details Page
       ↓
   useGame(id)
       ↓
gameService.getGame(id)
       ↓
apiClient.get(...)
       ↓
GET /api/games/:id
       ↓
Go Fiber
```

Reviews follow a similar flow:

```text
Game Details
       ↓
useReviews(gameId)
       ↓
reviewService.getReviews(gameId)
       ↓
apiClient.get(...)
       ↓
GET /api/games/:id/reviews
```

---

# 12. Data Flow — Create Review

The review submission flow should be:

```text
ReviewForm
     ↓
useReviews()
     ↓
reviewService.createReview()
     ↓
apiClient.post()
     ↓
POST /api/games/:id/reviews
     ↓
Go Fiber
     ↓
Created Review
     ↓
Hook updates local review state
     ↓
ReviewList re-renders
```

The user must see the new review immediately after a successful submission.

A full page reload should not be required.

---

# 13. State Management

Do not introduce a global state management library for this application.

Prefer:

- React state.
- Local component state.
- Custom hooks.
- Next.js Server Components where appropriate.

Use client-side state only where interaction requires it.

Examples of local state:

```text
Review form values
Submitting state
Validation errors
Review list after mutation
```

Do not move every piece of state into a global store.

---

# 14. Server vs Client Components

Use Next.js Server Components by default where appropriate.

Use `"use client"` only when the component requires client-side functionality such as:

- `useState`
- `useEffect`
- Event handlers
- Interactive forms
- Browser APIs

For example:

```text
app/games/[id]/page.tsx
        ↓
GameDetails
        ↓
ReviewList
        ↓
ReviewForm (Client Component)
```

Do not mark the entire application as `"use client"` without a reason.

---

# 15. Error Handling

Every asynchronous operation should have an appropriate error state.

Examples:

```text
Failed to load games.
Failed to load game details.
Failed to load reviews.
Failed to submit review.
```

Errors should be handled at the appropriate layer.

The API client should handle HTTP-level errors.

The hook should expose the error state.

The UI should decide how to present the error.

Example:

```text
API Client
    ↓
throws API error
    ↓
Hook catches/exposes error
    ↓
Component renders error state
```

---

# 16. Loading States

The UI should communicate when data is being loaded.

Examples:

```text
Games loading...
Game details loading...
Reviews loading...
Submitting review...
```

Avoid showing an empty state while data is still loading.

Distinguish between:

```text
Loading
Empty
Error
Success
```

These are different UI states.

---

# 17. Form Handling

Review form must validate:

```text
Reviewer name
Review text
Rating
```

Validation should happen before submitting to the backend for a good user experience.

However, frontend validation does NOT replace backend validation.

The backend remains the final authority.

Expected flow:

```text
Frontend validation
        ↓
API request
        ↓
Backend validation
        ↓
Create Review
```

---

# 18. Testing Strategy

Testing should focus on behavior rather than implementation details.

Priority:

```text
1. User-visible behavior
2. Hook behavior
3. Service behavior
4. Utility functions
```

Do not write tests solely to increase coverage.

---

# 19. Component Testing

Important components should have tests.

Examples:

```text
GameList
GameCard
ReviewList
ReviewCard
ReviewForm
```

Test behavior such as:

- Game appears in the list.
- Review appears.
- Form accepts input.
- Validation error appears.
- Submit button works.
- Submission state is displayed.
- New review appears after successful submission.
- API failure produces an error state.

Avoid testing implementation details such as:

```text
"setState was called"
```

unless there is a strong reason.

---

# 20. Hook Testing

Hooks should be tested when they contain meaningful behavior.

For example:

```text
useReviews()
```

may coordinate:

- Fetching reviews.
- Loading state.
- Error state.
- Creating a review.
- Updating the local review list.

Test the observable behavior rather than internal implementation.

---

# 21. Service Testing

Services should be independently testable.

For example:

```text
gameService.getGames()
reviewService.getReviews()
reviewService.createReview()
```

Tests should verify that the service correctly interacts with the API client and handles expected responses/errors.

Do not test the HTTP implementation again if that is already covered by the API client tests.

---

# 22. Avoid Duplicate Responsibility

Do not implement the same logic in multiple layers.

Bad:

```text
Component validates rating
Hook validates rating
Service validates rating
```

Instead:

```text
Component
  → basic user-input validation

Service
  → application operation

Backend
  → authoritative business validation
```

Keep responsibility clear.

---

# 23. No Over-Engineering

Do not introduce:

```text
Redux
Zustand
React Query
SWR
GraphQL
Next.js API Routes
Repository pattern
Dependency injection framework
Generic service factory
Generic API abstraction
```

unless a real requirement emerges.

For this application, the following is enough:

```text
UI
 ↓
Hook
 ↓
Service
 ↓
API Client
 ↓
Go Fiber
```

---

# 24. Definition of Done

Frontend work is considered complete when:

- [ ] UI behavior matches the acceptance criteria.
- [ ] Components are reasonably small.
- [ ] API calls are not scattered throughout UI components.
- [ ] Hooks manage relevant UI state.
- [ ] Services encapsulate application operations.
- [ ] API communication is centralized.
- [ ] Types are explicit.
- [ ] No unnecessary `any` is used.
- [ ] Loading states are handled.
- [ ] Error states are handled.
- [ ] Empty states are handled where appropriate.
- [ ] Form validation exists.
- [ ] Automated tests cover important behavior.
- [ ] Tests pass.
- [ ] Lint/type checks pass.
- [ ] Code is formatted.
- [ ] No unnecessary dependencies were introduced.

---

# 25. Architecture Summary

The final frontend architecture should remain understandable at a glance:

```text
                    ┌───────────────┐
                    │      UI       │
                    │   Components  │
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │     Hooks     │
                    │ UI State /    │
                    │ Orchestration │
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │   Services    │
                    │ Application   │
                    │  Operations   │
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │  API Client   │
                    │ HTTP / Routes │
                    └───────┬───────┘
                            │
                         REST API
                            │
                            ▼
                    ┌───────────────┐
                    │   Go Fiber    │
                    │    Backend    │
                    └───────────────┘
```

The architecture intentionally favors **clarity, separation of concerns, testability, and maintainability** over unnecessary abstraction.

The application is small, so the architecture should remain small as well.
