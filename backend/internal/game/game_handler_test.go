package game_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/game-review/backend/internal/container"
	"github.com/game-review/backend/internal/game"
	"github.com/game-review/backend/internal/server"
	"github.com/game-review/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func setupGameApp() *fiber.App {
	c := container.New()
	c.GameRepository.Seed([]game.Game{
		{ID: "1", Title: "Elden Ring", Genre: "Action RPG", Platform: "PC"},
		{ID: "2", Title: "The Witcher 3", Genre: "RPG", Platform: "PC"},
	})

	app := fiber.New()
	server.RegisterRoutes(app, c)
	return app
}

func TestHandlerGetGames(t *testing.T) {
	app := setupGameApp()

	req := httptest.NewRequest("GET", "/api/games", nil)
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

func TestHandlerGetGameByID(t *testing.T) {
	app := setupGameApp()

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

func TestHandlerGetGameByIDNotFound(t *testing.T) {
	app := setupGameApp()

	req := httptest.NewRequest("GET", "/api/games/nonexistent", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}
