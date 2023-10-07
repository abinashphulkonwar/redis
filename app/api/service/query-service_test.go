package service_test

import (
	"testing"

	"github.com/abinashphulkonwar/redis/api/service"
)

func TestStrimgHandler(t *testing.T) {
	command, err := service.GetCommands(" set sjdnkjsnakd sajkjnf ahjsbdjsad hjsabdjbadj akjhjsbdjasbdj                    jsdhkjnfkjnkjd\n      ")
	if err != nil {
		t.Error(err)
	}
	println(command.Command, command.Key, command.Value, command.EX, command.IF_NOT)
}
