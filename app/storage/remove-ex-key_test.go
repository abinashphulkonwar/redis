package storage_test

import (
	"testing"

	"github.com/abinashphulkonwar/redis/storage"
)

func TestExpirService(t *testing.T) {
	expireService := storage.InitExpireService()
	expireService.Start()
}
