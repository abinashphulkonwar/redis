package main

import (
	"github.com/abinashphulkonwar/redis/api"
	"github.com/abinashphulkonwar/redis/storage"
)

func main() {
	queue := storage.InitQueue()
	app := api.App(queue)
	go storage.DBCommandsHandler(queue)
	app.Listen(":3000")
}
