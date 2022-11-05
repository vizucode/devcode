package todoservice

import (
	todocore "devcode/domains/todo/core"
	"devcode/exceptions"
	"fmt"

	"gorm.io/gorm"
)

type todoService struct {
	repo todocore.IRepoTodo
}

func New(repo todocore.IRepoTodo) *todoService {
	return &todoService{
		repo: repo,
	}
}

func (r *todoService) Create(todoCore todocore.Core) todocore.Core {
	result, err := r.repo.Insert(todoCore)
	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	return result
}

func (r *todoService) Update(todoCore todocore.Core) todocore.Core {
	result, err := r.repo.Update(todoCore)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(exceptions.NewInternalServerError(err.Error()))
		} else {
			msg := fmt.Sprintf("Todo with ID %d Not Found", todoCore.Id)
			panic(exceptions.NewNotFoundError(msg))
		}
	}

	return result
}

func (r *todoService) Delete(todoCore todocore.Core) todocore.Core {
	result, err := r.repo.Delete(todoCore)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(exceptions.NewInternalServerError(err.Error()))
		} else {
			msg := fmt.Sprintf("Todo with ID %d Not Found", todoCore.Id)
			panic(exceptions.NewNotFoundError(msg))
		}
	}

	return result
}

func (r *todoService) FindAll(todoCore todocore.Core) []todocore.Core {
	result, err := r.repo.GetAll(todoCore)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(exceptions.NewInternalServerError(err.Error()))
		} else {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
	}

	return result
}

func (r *todoService) FindById(todoCore todocore.Core) todocore.Core {
	result, err := r.repo.GetById(todoCore)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(exceptions.NewInternalServerError(err.Error()))
		} else {
			msg := fmt.Sprintf("Todo with ID %d Not Found", todoCore.Id)
			panic(exceptions.NewNotFoundError(msg))
		}
	}

	return result
}
