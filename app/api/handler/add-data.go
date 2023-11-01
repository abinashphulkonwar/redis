package handler

import (
	"encoding/json"

	"github.com/abinashphulkonwar/redis/commands"
	"github.com/abinashphulkonwar/redis/internalstorage"
	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

func AddData(queue *internalstorage.Queue) func(c *fiber.Ctx) error {
	handler := func(c *fiber.Ctx) error {

		buf := c.Body()

		body := storage.RequestBody{}

		err := json.Unmarshal(buf, &body)

		if err != nil {
			return err
		}

		switch body.Commands {
		case commands.LSET:
			return InsertToQueue(c, &body, queue)
		case commands.LPUSH:
			return InsertToQueue(c, &body, queue)
		case commands.RPUSH:
			return InsertToQueue(c, &body, queue)
		case commands.TEXT:
			return InsertToQueue(c, &body, queue)
		case commands.NUMBER:
			intVal, isInt := body.Data.Value.(int)
			floatVal, isFloat := body.Data.Value.(float64)

			if !isInt && !isFloat {
				return fiber.NewError(fiber.StatusUnprocessableEntity, "value of 'type' 'number' key should be int")
			}
			if isInt {
				body.Data.Value = intVal
			}
			if isFloat {
				body.Data.Value = int(floatVal)

			}
			return InsertToQueue(c, &body, queue)
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"body":    body.Data,
			"message": "unvalid command",
		})
	}
	return handler
}

func InsertToQueue(c *fiber.Ctx, body *storage.RequestBody, queue *internalstorage.Queue) error {

	queue.Insert(&storage.DBCommands{
		Connection: c,
		Payload:    body,
	})
	return c.JSON(fiber.Map{
		"status": "error",
		"body":   body.Data,
	})

}
