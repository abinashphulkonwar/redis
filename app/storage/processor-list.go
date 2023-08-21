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
	var list *List = nil

	if data.Payload.IfNotExist && listRef != nil {
		println("data.Payload.IfNotExist")
		data.Payload.Data.Value = listRef.Value
		data.Payload.Data.Type = listRef.Type
		data.Payload.Data.EX = listRef.EX
		return
	}

	if isFound {
		if listRef.Type == commands.LSET {
			list = listRef.Value.(*List)
		} else {
			list = Init()
		}
	} else {
		list = Init()

	}
	AddValToList(data, list)
	data.Payload.Data.Value = list
	data.Payload.Data.Type = commands.LSET

}
