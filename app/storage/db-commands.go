package storage

import (
	"github.com/abinashphulkonwar/redis/commands"
	"github.com/gofiber/fiber/v2"
)

type DBCommands struct {
	Payload    *RequestBody
	Connection *fiber.Ctx
}

func DBCommandsHandler(queue *Queue) {
	for {

		val, isFound := queue.Get()
		if isFound {
			data := val.(*DBCommands)
			if data.Payload != nil {
				if data.Payload.Type == commands.LSET {
					ListProcessor(data)
				}
				Set(data.Payload.Key, &data.Payload.Data)
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
