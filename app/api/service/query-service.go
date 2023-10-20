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
	switch c.Command {
	case commands.GET:
		return true, ""
	case commands.LGET:
		return true, ""
	case commands.RGET:
		return true, ""
	}
	return false, "not a valid GET method"
}
func (c *Command) ValidateSet() (bool, string) {
	switch c.Command {
	case commands.SET:
		return true, ""
	}
	return false, "not a valid SET method"
}

func (c *Command) Validate() (bool, string) {
	if c.Key == "" {
		return false, "Key is undefined"
	}

	isGet, _ := c.ValidateGet()
	isSet, _ := c.ValidateGet()

	if !isGet && !isSet {
		return false, "Unvalid command"
	}
	return true, ""
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
