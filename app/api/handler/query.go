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

		isValid, message := query_struct.ValidateGet()

		if !isValid {
			return fiber.NewError(fiber.StatusUnprocessableEntity, message)
		}

		data, isFound := storage.Get(query_struct.Key)
		if !isFound {
			return fiber.NewError(fiber.StatusNotFound, "data not found")
		}

		if data.Type == storage.StringType || data.Type == storage.IntType {

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "Success",
				"data":   data.Value,
			})
		}
		if data.Type == storage.ListType {

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "Success_List_Type",
				"data":   "List Type",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Error",
			"data":   nil,
		})
	}
	return handler
}
