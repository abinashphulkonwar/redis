package handler

import (
	"strings"

	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

func GetQuery(queue *storage.Queue) func(c *fiber.Ctx) error {
	handler := func(c *fiber.Ctx) error {

		query := c.Query("command")

		query = strings.Trim(query, " \n\t\r")
		if len(query) == 0 {
			fiber.NewError(fiber.StatusNoContent, "no command found")
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"body":    query,
			"message": "unvalid command",
		})
	}
	return handler
}
