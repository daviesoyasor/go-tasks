package controllers

import (
	"github.com/daviesoyasor/go-tasks/pkg/configs"
	"github.com/gofiber/fiber/v2"
)

func GetSheetsData(c *fiber.Ctx) error {
	srv, err := configs.GetGoogleClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Unable to retrieve Sheets client",
		})
	}

	spreadsheetID := "1fsZRUHaQWI0TWQUE7fykDn37JnMl9gtjZD1bzqY4ww8"
	sheetName := "Sheet1"
	rangeName := "A2:D"
	values, err := configs.ReadRange(srv, spreadsheetID, sheetName, rangeName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Range incorrect",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  values,
	})
}
