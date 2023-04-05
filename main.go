package main

import (
	"github.com/daviesoyasor/go-tasks/pkg/configs"
	"github.com/daviesoyasor/go-tasks/pkg/middleware"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Define a new Fiber app with config.
	// Function is used to create and configure a new instance of the fiber.App struct 
	app := configs.InitializeNewFiberInstance()

	// Fiber Middlewares (APPLICATION LEVEL MIDDLEWARE)
	middleware.FiberMiddleware(app)

	// Routes

	// Start server (with graceful shutdown).
	configs.StartServerWithGracefulShutdown(app)
}
