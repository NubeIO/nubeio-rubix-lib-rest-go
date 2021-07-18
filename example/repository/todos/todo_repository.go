package todos

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/example/model/todos"
)

var (
	idSeq    uint = 0
	todoList      = []*todos.TODO{}
)

func GetTODOs() []*todos.TODO {
	return todoList
}

func GetTODO(id uint) *todos.TODO {
	for _, todo := range todoList {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func CreateTODO(dto TODODto) *todos.TODO {
	idSeq += 1

	todo := &todos.TODO{
		ID: idSeq,
		Title: dto.Title,
		IsChecked: dto.IsChecked,
	}

	todoList = append(todoList, todo)
	return todo
}

func UpdateTODO(id uint, dto TODODto) error {
	for _, todo := range todoList {
		if todo.ID == id {
			todo.Title = dto.Title
			todo.IsChecked = dto.IsChecked
			return nil
		}
	}
	return errors.New("todo not found")
}

func DeleteTODO(id uint) error {
	for i, todo := range todoList {
		if todo.ID == id {
			todoList[i] = todoList[len(todoList)-1]
			todoList = todoList[:len(todoList)-1]
			return nil
		}
	}
	return errors.New("todo not found")
}