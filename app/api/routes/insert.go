package routes

import (
	"github.com/abinashphulkonwar/redis/api/handler"
	"github.com/abinashphulkonwar/redis/internalstorage"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(queue *internalstorage.Queue) func(router fiber.Router) {
	router := func(router fiber.Router) {
		router.Post("/add", handler.AddData(queue))
	}
	return router
}
