package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"goproj/configs"
	"goproj/middleware"
	"goproj/routes"
	"goproj/utils/schedule"
	"log"
)

// @title Swagger API
// @description This is a sample swagger for Fiber
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	schedule.ProcessScheduler()

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)
	routes.Router(app)

	log.Fatal(app.Listen(":3000"))
}
