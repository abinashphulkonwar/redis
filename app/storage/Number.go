package storage

import "github.com/abinashphulkonwar/redis/commands"

func NumberProcessor(data *DBCommands) {
	val, isFound := Get(data.Payload.Key)
	data.Payload.Data.Type = commands.TYPE_NUMBER

	if isFound && data.Payload.IfNotExist {
		return
	} else if isFound && val.Type == commands.TYPE_NUMBER {
		println("data inc")
		data.Payload.Data.Value = val.Value.(int) + data.Payload.Data.Value.(int)
		data.Payload.Data.Type = val.Type
		data.Payload.Data.EX = val.EX
		Set(data.Payload.Key, &data.Payload.Data)
	} else {
		Set(data.Payload.Key, &data.Payload.Data)
		println("data add defult")

	}

}
