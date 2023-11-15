package storage

import (
	"sync"
)

type Data struct {
	Value interface{}
	EX    int
	Type  string
}
type RequestBody struct {
	Data       Data
	Key        string
	Commands   string
	IfNotExist bool
}

var hMap = sync.Map{}

func Get(key string) (*Data, bool) {

	data, isFound := hMap.Load(key)
	if !isFound {
		return nil, false
	}
	res := data.(*Data)
	if res == nil {
		return nil, false
	}
	return res, true
}

func Set(key string, data *Data) {
	hMap.Store(key, data)

}

func Remove(Key string) {
	hMap.Delete(Key)
}
