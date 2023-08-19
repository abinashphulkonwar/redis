package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/abinashphulkonwar/redis/api"
	"github.com/abinashphulkonwar/redis/api/handler"
	"github.com/abinashphulkonwar/redis/storage"
)

func TestApp(t *testing.T) {
	queue := storage.InitQueue()
	go storage.DBCommandsHandler(queue)
	body := storage.RequestBody{
		Key: "id",
		Data: storage.Data{
			Value: "🚀",
			EX:    100000,
		},
		Type:       handler.LIST,
		Commands:   handler.C_RPUSH,
		IfNotExist: true,
	}
	data, err := json.Marshal(body)

	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest("POST", "/api/write/add", bytes.NewReader(data))

	app := api.App(queue)
	resp, err := app.Test(req)
	if err != nil {
		t.Errorf("Error adding key value pair " + err.Error())
		return
	}

	res, _ := io.ReadAll(resp.Body)
	println(resp.StatusCode, string(res))
	stored_data, isFound := storage.Get("id")
	if isFound {
		switch v := stored_data.(type) {
		case string:
			fmt.Println(string(res), v)
		default:
			fmt.Println("Stored data is not a string")
		}
	}
}
