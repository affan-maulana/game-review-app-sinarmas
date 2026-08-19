package tests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/game-review/backend/internal/container"
	"github.com/game-review/backend/internal/game"
	"github.com/game-review/backend/internal/review"
	"github.com/game-review/backend/internal/server"
	"github.com/game-review/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	c := container.New()
	c.GameRepository.Seed([]game.Game{
		{ID: "1", Title: "Elden Ring", Genre: "Action RPG", Platform: "PC"},
		{ID: "2", Title: "Hollow Knight", Genre: "Metroidvania", Platform: "Switch"},
	})
	c.ReviewRepository.Seed([]review.Review{
		{ID: "r1", GameID: "1", ReviewerName: "Alex", Text: "Great!", Rating: 5},
	})

	app := fiber.New()
	server.RegisterRoutes(app, c)
	return app
}

func TestHealthCheck(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var apiResp response.APIResponse
	json.Unmarshal(body, &apiResp)

	if apiResp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", apiResp.Status)
	}
}

func TestGetGames(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/games/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var apiResp response.APIResponse
	json.Unmarshal(body, &apiResp)

	if apiResp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", apiResp.Status)
	}

	if apiResp.Data == nil {
		t.Fatal("expected non-nil data in response")
	}
}

func TestGetGameByID(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/games/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var apiResp response.APIResponse
	json.Unmarshal(body, &apiResp)

	if apiResp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", apiResp.Status)
	}
}

func TestGetGameByIDNotFound(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/games/nonexistent", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestGetReviews(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/games/1/reviews", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var apiResp response.APIResponse
	json.Unmarshal(body, &apiResp)

	if apiResp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", apiResp.Status)
	}
}

func TestGetReviewsGameNotFound(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/games/nonexistent/reviews", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestCreateReview(t *testing.T) {
	app := setupApp()

	body, _ := json.Marshal(review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Amazing game!",
		Rating:       5,
	})

	req := httptest.NewRequest("POST", "/api/games/1/reviews", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	var apiResp response.APIResponse
	json.Unmarshal(respBody, &apiResp)

	if apiResp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", apiResp.Status)
	}
}

func TestCreateReviewGameNotFound(t *testing.T) {
	app := setupApp()

	body, _ := json.Marshal(review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Amazing game!",
		Rating:       5,
	})

	req := httptest.NewRequest("POST", "/api/games/nonexistent/reviews", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestCreateReviewInvalidBody(t *testing.T) {
	app := setupApp()

	body, _ := json.Marshal(review.CreateReviewRequest{
		ReviewerName: "",
		Text:         "Great!",
		Rating:       5,
	})

	req := httptest.NewRequest("POST", "/api/games/1/reviews", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestCreateReviewInvalidRating(t *testing.T) {
	app := setupApp()

	body, _ := json.Marshal(review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Great!",
		Rating:       10,
	})

	req := httptest.NewRequest("POST", "/api/games/1/reviews", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestNonExistentRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/nonexistent", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}
