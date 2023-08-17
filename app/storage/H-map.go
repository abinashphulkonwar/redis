package storage

import "sync"

type Data struct {
	Type       string
	Value      interface{}
	EX         int
	IfNotExist bool
}

var hMap = sync.Map{}

func Get(key string) (*Data, bool) {
	data, isFound := hMap.Load(key)
	return data.(*Data), isFound
}

func Set(key string, data *Data) {
	hMap.Store(key, data)
}
