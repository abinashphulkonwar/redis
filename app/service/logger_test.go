package service_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/abinashphulkonwar/redis/service"
	"github.com/brianvoe/gofakeit/v6"
)

func TestLoggerWrite(t *testing.T) {
	os.Setenv("mode", "Test")
	logger := service.InitLogger("log")

	logger.Add(&service.Log{
		Time:    gofakeit.Date().String(),
		Status:  strconv.Itoa(gofakeit.HTTPStatusCode()),
		Path:    gofakeit.URL(),
		Method:  gofakeit.HTTPMethod(),
		Command: gofakeit.Username(),
		Key:     gofakeit.RandomString([]string{gofakeit.Name(), gofakeit.Email()}),
		Value:   gofakeit.RandomString([]string{gofakeit.Name(), gofakeit.Email(), gofakeit.Username()}),
	})
	logger.New()
}
func TestLoggerRead(t *testing.T) {
	os.Setenv("mode", "Test")
	logger := service.InitLogger("log")
	logger.ReadLogs()
}
