package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/pablon/backend-pablon/docs" // caminho do módulo gerado por swag
	"github.com/pablon/backend-pablon/config"
	"github.com/pablon/backend-pablon/routes"
)

// @title Pablon API
// @version 1.0
// @description API para autenticação e parceiros
// @host localhost:3000
// @BasePath /api
func main() {
	cfg := config.GetConfig()
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // swagger docs

	routes.SetupRoutes(app) // define as rotas

	port := fmt.Sprintf(":%s", cfg.Port)
	app.Listen(port)
}
