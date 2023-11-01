package storage

import (
	"time"

	"github.com/abinashphulkonwar/redis/commands"
	"github.com/abinashphulkonwar/redis/internalstorage"
	"github.com/abinashphulkonwar/redis/service"
	"github.com/gofiber/fiber/v2"
)

type DBCommands struct {
	Payload    *RequestBody
	Connection *fiber.Ctx
}

func DBCommandsHandler(queue *internalstorage.Queue, logger *service.Logger) {

	for {

		val, isFound := queue.Get()
		if isFound {
			data := val.(*DBCommands)
			if data.Payload != nil {

				switch data.Payload.Commands {
				case commands.LSET:
					ListProcessor(data)
				case commands.RPUSH:
					ListProcessor(data)
				case commands.LPUSH:
					ListProcessor(data)
				case commands.TEXT:
					TextProcessor(data)
				case commands.NUMBER:
					NumberProcessor(data)
				}
				logger.Add(&service.Log{
					Time:    time.Now().String(),
					Status:  "success",
					Path:    data.Connection.Path(),
					Method:  data.Connection.Method(),
					Command: data.Payload.Commands,
					Key:     data.Payload.Key,
					Value:   data.Payload.Data.Value.(string),
				})
			}
			queue.Remove()

			// if data.Connection != nil {
			// 	println("🚀", data.Connection.BaseURL())
			// 	data.Connection.JSON(fiber.Map{
			// 		"status": "done",
			// 		"value":  data.Payload.Key,
			// 	})
			// }

		}
	}
}
