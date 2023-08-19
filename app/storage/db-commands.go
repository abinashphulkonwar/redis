package storage

import "github.com/gofiber/fiber/v2"

type DBCommands struct {
	Payload    *RequestBody
	Connection *fiber.Ctx
}

func DBCommandsHandler(queue *Queue) {
	for {

		val, isFound := queue.Get()
		if isFound {
			data := val.(DBCommands)
			if data.Payload != nil {
				Set(data.Payload.Key, &data.Payload.Data)
			}
			queue.Remove()

			// if data.Connection != nil {
			// 	println("🚀")
			// 	data.Connection.JSON(fiber.Map{
			// 		"status": "done",
			// 	})
			// }

		}
	}
}
