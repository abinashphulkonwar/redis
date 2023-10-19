package storage

import "github.com/abinashphulkonwar/redis/commands"

func TextProcessor(data *DBCommands) {

	_, isFound := Get(data.Payload.Key)
	if isFound && data.Payload.IfNotExist {
		return
	}
	data.Payload.Data.Type = commands.TYPE_STRING
	Set(data.Payload.Key, &data.Payload.Data)

}
