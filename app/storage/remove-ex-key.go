package storage

import (
	"time"
)

type ExpireService struct {
}

func InitExpireService() *ExpireService {
	return &ExpireService{}
}

func (ex *ExpireService) Start() {
	for {
		println("expiration service")

		curren_time := time.Now()
		hMap.Range(func(key, value any) bool {
			val := value.(Data)
			node_time := time.Now()
			node_time = node_time.Add(time.Second * time.Duration(val.EX))
			if curren_time.Compare(node_time) == 1 || curren_time.Compare(node_time) == 0 {
				Remove(key.(string))
			}
			return true
		})

		time.Sleep(time.Minute * time.Duration(1))

	}
}
