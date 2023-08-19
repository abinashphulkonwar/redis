package storage

import (
	"sync"
)

type Data struct {
	Value interface{}
	EX    int
}
type RequestBody struct {
	Data       Data
	Key        string
	Type       string
	Commands   string
	IfNotExist bool
}

var hMap = sync.Map{}

func Get(key string) (interface{}, bool) {
	data, isFound := hMap.Load(key)
	if !isFound {
		return nil, false
	}
	res := data.(*Data)
	if res == nil {
		return nil, false
	}
	return res.Value, true
}

func Set(key string, data *Data) {
	hMap.Store(key, data)

}
