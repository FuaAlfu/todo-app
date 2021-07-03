package usecases

import "github.com/FuaAlfu/todo-app/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error){
	return nil, ErrInternal
}