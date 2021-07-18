package main

import (
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/example/controller/todos"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
)

func main() {
	app := rest.New(3)

	app.Controller(todos.New())

	err := app.Run(":8080")
	if err != nil {
		return 
	}


}
