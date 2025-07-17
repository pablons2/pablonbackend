package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pablon/backend-pablon/handlers"
	"github.com/pablon/backend-pablon/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/google", handlers.GoogleLogin)

	secured := api.Group("/secure", middleware.Protected())
	secured.Get("/profile", handlers.GetProfile)
}
