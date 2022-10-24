package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"goproj/internal/result"
	"goproj/utils"
	"strings"
	"time"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodDelete,
				fiber.MethodPut,
			}, ","),
			AllowCredentials: true,
		}),
		logger.New(),
	)
}

func TokenAuthMiddleware(c *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		if err.Error() == "Token is expired" {
			return c.Status(fiber.StatusUnauthorized).JSON(result.Result[interface{}]{
				Success:      false,
				ErrorMessage: "토큰이 만료됐습니다.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(result.Result[interface{}]{
			Success:      false,
			ErrorMessage: err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(result.Result[interface{}]{
			Success:      false,
			ErrorMessage: "토큰이 만료됐습니다.",
		})
	}

	return c.Next()
}
