package handler

import (
	"strings"

	"github.com/abinashphulkonwar/redis/api/service"
	"github.com/abinashphulkonwar/redis/commands"
	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

func listError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success_List_Error",
		"data":    nil,
		"message": message,
	})
}

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

		if data.Type == commands.TYPE_STRING || data.Type == commands.TYPE_NUMBER {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "Success",
				"data":   data.Value,
			})
		} else if data.Type == commands.TYPE_LIST {
			var list *storage.List
			var val interface{}
			var status uint8
			switch query_struct.Command {
			case commands.LGET:
				list = data.Value.(*storage.List)
				val, status = list.LGet()
				if status == 0 {
					return listError(c, "not found")
				}
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"status": "Success_List_Type",
					"data":   val,
				})
			case commands.RGET:
				list = data.Value.(*storage.List)
				val, status = list.RGet()
				if status == 0 {
					return listError(c, "not found")
				}
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"status": "Success_List_Type",
					"data":   val,
				})

			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Success_List_Type",
				"data":    "List Type",
				"Is_LIST": true,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Error",
			"data":   nil,
		})
	}
	return handler
}
