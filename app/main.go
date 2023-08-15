package main

import (
	"github.com/abinashphulkonwar/redis/api"
)

func main() {
	app := api.App()
	app.Listen(":3000")
}
