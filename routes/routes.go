package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/daviesoyasor/go-tasks/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func Routes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1") 

	// Routes for GET method:
	route.Get("/sheets", controllers.GetSheetsData)     

}
