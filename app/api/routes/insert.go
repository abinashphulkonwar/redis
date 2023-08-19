package routes

import (
	"github.com/abinashphulkonwar/redis/api/handler"
	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(queue *storage.Queue) func(router fiber.Router) {
	router := func(router fiber.Router) {
		router.Post("/add", handler.AddData(queue))
	}
	return router
}
