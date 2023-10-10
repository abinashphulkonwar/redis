package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/abinashphulkonwar/redis/commands"
)

type Command struct {
	Command string
	Key     string
	Value   string
	EX      int
	IF_NOT  bool
}

func (c *Command) ValidateGet() (bool, string) {
	if c.Command != commands.GET {
		return false, "Not a valid command"
	}
	if c.Key == "" {
		return false, "Key is undefined"
	}
	return true, ""
}
func (c *Command) Validate() (bool, string) {

	switch c.Command {
	case commands.SET:
		return true, ""
	case commands.GET:
		return c.ValidateGet()
	}

	return false, "Unvalid command"
}

func GetCommands(arg string) (*Command, error) {
	command := Command{}
	arg = strings.Trim(arg, " \n\t\r") + " "

	temp := ""
	state := 0
	isEx := false
	isIFNot := false
	for _, v := range arg {
		current := string(v)
		if state == 0 && current == " " {
			command.Command = strings.ToUpper(temp)
			temp = ""
			state++

		} else if state == 1 && current == " " {
			command.Key = temp
			temp = ""
			state++

		} else if state == 2 && current == " " {
			if temp == "EX" {
				isEx = true
				state++
			} else if temp == "IF_NOT" {
				isIFNot = true
				state++
			} else {
				command.Value = command.Value + temp + current
			}
			temp = ""
		} else if (state == 3 || state == 4) && current == " " {

			if isEx {
				isEx = false
				i, err := strconv.Atoi(temp)
				if err != nil {
					return nil, err
				}
				command.EX = i
				state++

			} else if isIFNot {
				isIFNot = false
				state++
				if temp == "true" {
					command.IF_NOT = true
				} else {
					command.IF_NOT = false
				}

			}
			temp = ""
		}

		if current != " " {
			temp = temp + current
		}

	}

	isValid, message := command.Validate()

	if isValid {
		return &command, nil
	} else {
		return nil, errors.New(message)
	}
}
