package service_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/redis/service"
)

func TestLoggerWrite(t *testing.T) {
	os.Setenv("mode", "Test")
	logger := service.InitLogger("log")
	logger.Add(&service.Log{
		Time:    "sjkjsd",
		Status:  "1",
		Path:    "string",
		Method:  "stringjabsjh",
		Command: "string",
		Key:     "string",
		Value:   "string\najbjabd",
	})
	logger.New()
}
func TestLoggerRead(t *testing.T) {
	os.Setenv("mode", "Test")
	logger := service.InitLogger("log")
	logger.Read()
}
