package todos

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/example/repository/todos"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/example/response"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func New() rest.IController {
	c := rest.Controller("todos")

	c.GET("/", getTODOs)
	c.POST("/", createTODO)
	c.SUB("/id")
	c.GET("/", getTODO)
	c.PUT("/", updateTODO)
	c.PATCH("/", updateTODO)
	c.DELETE("/", deleteTODO)


	return c
}

func getTODOs(*gin.Context) rest.IResponse {
	return response.Data(todos.GetTODOs())
}

func getTODO(ctx *gin.Context) rest.IResponse {
	id, err := resolveID(ctx)
	if err != nil {
		return response.BadEntity("ID must be valid")
	}

	todo := todos.GetTODO(id)
	if todo == nil {
		return response.NotFound("todo not found")
	}

	return response.Data(todo)
}

func createTODO(ctx *gin.Context) rest.IResponse {
	dto, err := getDTO(ctx)
	if err != nil {
		return response.BadEntity(err.Error())
	}

	return response.Created(todos.CreateTODO(dto).ID)
}

func updateTODO(ctx *gin.Context) rest.IResponse {
	fmt.Println("da")
	id, err := resolveID(ctx)
	if err != nil {
		return response.BadEntity("ID must be valid")
	}
	fmt.Println("da")

	dto, err := getDTO(ctx)
	if err != nil {
		return response.BadEntity(err.Error())
	}
	fmt.Println("da")

	if err := todos.UpdateTODO(id, dto); err != nil {
		return response.NotFound(err.Error())
	}
	fmt.Println("da")

	return response.Data(todos.GetTODO(id))
}

func deleteTODO(ctx *gin.Context) rest.IResponse {
	id, err := resolveID(ctx)
	if err != nil {
		return response.BadEntity("ID must be valid")
	}

	if err := todos.DeleteTODO(id); err != nil {
		return response.NotFound(err.Error())
	}

	return response.OK(id)

}

func getDTO(ctx *gin.Context) (dto todos.TODODto, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func resolveID(ctx *gin.Context) (uint, error) {
	id := ctx.Query("id")
	val, err := strconv.ParseUint(id, 10, 32)
	return uint(val), err
}