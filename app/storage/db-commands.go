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

				switch data.Payload.Type {
				case commands.LSET:
					ListProcessor(data)
					Set(data.Payload.Key, &data.Payload.Data)
				case commands.TEXT:
					TextHandler(data)
				}
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

func TextHandler(data *DBCommands) {

	_, isFound := Get(data.Payload.Key)

	if (!isFound && data.Payload.IfNotExist) || (isFound && !data.Payload.IfNotExist) {
		data.Payload.Data.Type = data.Payload.Type
		Set(data.Payload.Key, &data.Payload.Data)
	}

}
