package handler_test

import (
	"testing"

	"github.com/abinashphulkonwar/redis/api/handler"
)

func TestStrimgHandler(t *testing.T) {
	command, err := handler.GetCommands(" set sjdnkjsnakd sajkjnf ahjsbdjsad hjsabdjbadj akjhjsbdjasbdj                    jsdhkjnfkjnkjd\n      ")
	if err != nil {
		t.Error(err)
	}
	println(command.Command, command.Key, command.Value, command.EX, command.IF_NOT)
}
