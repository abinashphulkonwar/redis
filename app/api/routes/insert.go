package routes

import (
	"github.com/abinashphulkonwar/redis/api/handler"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(router fiber.Router) {
	router.Post("/add", handler.AddData)

}
