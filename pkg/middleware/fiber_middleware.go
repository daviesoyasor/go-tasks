package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),

		// Add simple logger.
		logger.New(),

		// Rate limit: 3 requests per 10 seconds max
		limiter.New(limiter.Config{
			Expiration: 10 * time.Second,
			Max:        3,
		}),

		NotFoundRoute,
	)
}
