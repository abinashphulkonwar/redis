package storage

import (
	"github.com/abinashphulkonwar/redis/commands"
)

func ListProcessor(data *DBCommands) {

	if data.Payload.Type == commands.LSET {
		list := Init()
		data.Payload.Data.Value = list
	} else if data.Payload.Type == commands.C_LPUSH || data.Payload.Type == commands.C_RPUSH {
		listRef, isFound := Get(data.Payload.Key)
		if isFound {
			if listRef.Type == commands.LSET {
				list := listRef.Value.(*List)
				if data.Payload.Type == commands.C_LPUSH {
					list.LPush(data.Payload.Data.Value)
				} else {
					list.RPush(data.Payload.Data.Value)
				}
			} else {
				list := Init()
				data.Payload.Data.Value = list
			}
		}
	}

}
