package handler_test

import (
	"strings"
	"testing"
)

func TestStrimgHandler(t *testing.T) {
	val := strings.Trim(" sjdnkjsnakd sajkjnf\n      ", " \n\t\r")
	println(val, "hjshj")
	commands := strings.Split(val, " ")
	println(commands[0], commands[1])
}
