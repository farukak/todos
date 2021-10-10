package route

import (
	"github.com/farukak/todos/handlers"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	api := app.Group("/api/v1")
	api.Get("/health", handlers.Health)

}
