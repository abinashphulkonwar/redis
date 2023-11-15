package main

import (
	"github.com/abinashphulkonwar/redis/api"
	"github.com/abinashphulkonwar/redis/internalstorage"
	"github.com/abinashphulkonwar/redis/service"
	"github.com/abinashphulkonwar/redis/storage"
)

func main() {
	queue := internalstorage.InitQueue()
	app := api.App(queue)
	loger := service.InitLogger("log")
	go loger.New()
	go storage.DBCommandsHandler(queue, loger)
	expireService := storage.InitExpireService()
	go expireService.Start()
	app.Listen(":3000")
}
