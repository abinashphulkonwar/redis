package handler

import "github.com/gofiber/fiber/v2"

func AddData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
