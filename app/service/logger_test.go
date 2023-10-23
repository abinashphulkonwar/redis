package service_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/redis/service"
)

func TestLogger(t *testing.T) {
	os.Setenv("mode", "Test")
	logger := service.InitLogger("log.txt")
	logger.Add(&service.Log{
		Time:    "sjkjsd",
		Status:  "1",
		Path:    "string",
		Method:  "string\njabsjh",
		Command: "string",
		Key:     "string",
		Value:   "string",
	})
	logger.New()
}
