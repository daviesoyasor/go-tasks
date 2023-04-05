package middleware

import "github.com/gofiber/fiber/v2"

// NotFoundRoute func for describe 404 Error route.
// [c *fiber.Ctx] represents the context of an HTTP request and response object.
// Just like in Node Js where you have (req, res)
func NotFoundRoute(c *fiber.Ctx) error {
	// Return HTTP 404 status and JSON response.
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": true,
		"msg":   "sorry, endpoint is not found",
	})
}
