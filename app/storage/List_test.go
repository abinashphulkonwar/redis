package storage_test

import (
	"testing"

	"github.com/abinashphulkonwar/redis/storage"
)

func TestList(t *testing.T) {
	list := storage.Init()

	list.RPush([]byte("data"))
	list.LPush([]byte("hiii"))
	list.RPush([]byte("1"))
	list.LPush(748743)
	list.RPush("🚀")
	list.RRemove()
	list.LRemove()

	list.Travers()
}
