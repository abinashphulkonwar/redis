package storage

func NumberProcessor(data *DBCommands) {
	val, isFound := Get(data.Payload.Key)
	data.Payload.Data.Type = data.Payload.Type
	if (!isFound && data.Payload.IfNotExist) || (isFound && !data.Payload.IfNotExist) {
		Set(data.Payload.Key, &data.Payload.Data)
	} else {
		data.Payload.Data.Value = val.Value.(int) + data.Payload.Data.Value.(int)
		data.Payload.Data.Type = val.Type
		data.Payload.Data.EX = val.EX
		Set(data.Payload.Key, &data.Payload.Data)
	}
}
