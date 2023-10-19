package storage

import (
	"github.com/abinashphulkonwar/redis/commands"
)

func AddValToList(data *DBCommands, list *List) {
	if data.Payload.Commands == commands.C_LPUSH {
		println("list LPUSH")
		list.LPush(data.Payload.Data.Value)

	} else {
		println("list RPUSH")
		list.RPush(data.Payload.Data.Value)
	}
}

func ListProcessor(data *DBCommands) {

	listRef, isFound := Get(data.Payload.Key)

	if isFound && data.Payload.IfNotExist {
		return
	}

	if data.Payload.IfNotExist && listRef != nil {
		println("data.Payload.IfNotExist")
		data.Payload.Data.Value = listRef.Value
		data.Payload.Data.Type = commands.TYPE_LIST
		data.Payload.Data.EX = listRef.EX
		Set(data.Payload.Key, &data.Payload.Data)
		return
	}

	var list *List = nil

	if isFound && listRef.Type == commands.TYPE_LIST {
		list = listRef.Value.(*List)
	} else {
		list = Init()

	}

	AddValToList(data, list)
	data.Payload.Data.Value = list
	data.Payload.Data.Type = commands.LSET
	Set(data.Payload.Key, &data.Payload.Data)

}
