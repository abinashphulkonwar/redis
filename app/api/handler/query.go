package handler

import (
	"strconv"
	"strings"

	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

type Command struct {
	Command string
	Key     string
	Value   string
	EX      int
	IF_NOT  bool
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
			command.Command = temp
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

	return &command, nil
}

func GetQuery(queue *storage.Queue) func(c *fiber.Ctx) error {
	handler := func(c *fiber.Ctx) error {

		query := c.Query("command")

		query = strings.Trim(query, " \n\t\r")
		if len(query) == 0 {
			fiber.NewError(fiber.StatusNoContent, "no command found")
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"body":    query,
			"message": "unvalid command",
		})
	}
	return handler
}
