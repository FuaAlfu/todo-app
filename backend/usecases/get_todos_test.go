package usecases_test

import (
	"fmt"
	"testing"

	"github.com/FuaAlfu/todo-app/backend/entities"
	"github.com/FuaAlfu/todo-app/backend/usecases"
)

type MockTodosRepo struct {
}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("missing..., orsomthing went wrong")
}

func TestGetTdos(t *Testing.T) {
	repo := new(MockTodosRepo)
	todos, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("expected an ErrInternal: Got: %v", err)
	}

	if todos != nil {
		t.Fatalf("expected todos to be nil: Got: %v", todos)
	}
}
