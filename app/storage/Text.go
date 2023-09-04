package storage

func TextProcessor(data *DBCommands) {

	_, isFound := Get(data.Payload.Key)

	if (!isFound && data.Payload.IfNotExist) || (isFound && !data.Payload.IfNotExist) {
		data.Payload.Data.Type = data.Payload.Type
		Set(data.Payload.Key, &data.Payload.Data)
	}

}
