package handler

import (
	"strings"

	"github.com/abinashphulkonwar/redis/api/service"
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

		query_struct, err := service.GetCommands(query)
		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Success",
			"query":  query_struct,
		})
	}
	return handler
}
