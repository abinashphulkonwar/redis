package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/abinashphulkonwar/redis/api"
	"github.com/abinashphulkonwar/redis/commands"
	"github.com/abinashphulkonwar/redis/storage"
)

func GetJson(body storage.RequestBody) []byte {

	data, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	return data
}

func TestApp(t *testing.T) {
	queue := storage.InitQueue()
	go storage.DBCommandsHandler(queue)
	body := storage.RequestBody{
		Key: "id",
		Data: storage.Data{
			Value: "🚀",
			EX:    100000,
		},
		Type:       commands.LSET,
		Commands:   commands.C_LPUSH,
		IfNotExist: true,
	}
	data := GetJson(body)
	req := httptest.NewRequest("POST", "/api/write/add", bytes.NewReader(data))
	body.Data.Value = "data"
	body.IfNotExist = false
	req1 := httptest.NewRequest("POST", "/api/write/add", bytes.NewReader(GetJson(body)))
	body.Data.Value = "add new data"
	body.Commands = commands.C_RPUSH
	body.IfNotExist = true
	req2 := httptest.NewRequest("POST", "/api/write/add", bytes.NewReader(GetJson(body)))

	app := api.App(queue)
	resp, err := app.Test(req)
	_, _ = app.Test(req1)
	_, _ = app.Test(req2)
	if err != nil {
		t.Errorf("Error adding key value pair " + err.Error())
		return
	}

	//res, _ := io.ReadAll(resp.Body)
	println(resp.StatusCode)
	stored_data, isFound := storage.Get("id")
	if isFound {
		switch data := stored_data.Value.(type) {
		case *storage.List:
			println(stored_data.Type)

			data.Travers()
		default:
			fmt.Println("Stored data is not a string")
		}
	}
}
