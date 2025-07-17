package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pablon/backend-pablon/utils"
)

// GoogleLogin godoc
// @Summary Valida token do Google recebido do NextAuth
// @Accept json
// @Produce json
// @Param token body map[string]string true "Token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/auth/google [post]
func GoogleLogin(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Corpo inválido"})
	}

	// Aqui você pode usar a Google API para validar o token (ou deixar pra frente)
	// Se quiser validar agora, use: https://oauth2.googleapis.com/tokeninfo?id_token=...

	token, err := utils.GenerateJWT("user@email.com") // substitua pelo email real extraído do token Google
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Falha ao gerar JWT"})
	}

	return c.JSON(fiber.Map{"token": token})
}
