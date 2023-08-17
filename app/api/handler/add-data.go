package handler

import (
	"encoding/json"

	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	LIST   = "LSET"
	TEXT   = "SET"
	NUMBER = "SET"
)

func AddData(c *fiber.Ctx) error {
	buf := c.Body()

	body := storage.Data{}

	err := json.Unmarshal(buf, &body)

	if err != nil {
		return err
	}

	if body.Type == LIST {
		return ListHandler(c, &body)
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func ListHandler(c *fiber.Ctx, body *storage.Data) error {

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
