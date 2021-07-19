package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/example/controller/todos"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"log"
)

func main() {
	app := rest.New(3)

	app.Controller(todos.New())
	children, err := rest.WithChildren("true")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(children, err)
	err = app.Run(":8080");if err != nil {
		log.Println(err)
		return
	}


}
