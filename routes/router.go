package routes

import (
	"goproj/controller"
	_ "goproj/docs"
	"goproj/middleware"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  true,
		DocExpansion: "none",
	}))

	Routes(app)
	TokenRoutes(app)
}

func Routes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/user", controller.GetUser)
	route.Post("/user/auth", controller.Authenticate)
}

func TokenRoutes(a *fiber.App) {
	route := a.Group("/api/admin", middleware.TokenAuthMiddleware)

	route.Get("/users", controller.GetUsers)
}
