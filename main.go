package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/daviesoyasor/go-tasks/pkg/middleware"
	"github.com/daviesoyasor/go-tasks/pkg/configs"
	
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Fiber Middlewares (APPLICATION LEVEL MIDDLEWARE)
	middleware.FiberMiddleware(app)

	// Routes

	// Start server (with graceful shutdown).
	configs.StartServerWithGracefulShutdown(app)
}
