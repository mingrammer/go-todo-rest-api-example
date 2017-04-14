package main

import (
	"github.com/mingrammer/go-todo-rest-api-example/app"
	"github.com/mingrammer/go-todo-rest-api-example/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
