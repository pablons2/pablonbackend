package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile information
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/secure/profile [get]
func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return c.JSON(fiber.Map{
		"email":   email,
		"message": "Bem-vindo à rota protegida!",
	})
}
