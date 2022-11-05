package todocore

type IRepoTodo interface {
	Insert(todoCore Core) (Core, error)
	Update(todoCore Core) (Core, error)
	Delete(todoCore Core) (Core, error)
	GetAll(todoCore Core) ([]Core, error)
	GetById(todoCore Core) (Core, error)
}

type IServiceTodo interface {
	Create(todoCore Core) Core
	Update(todoCore Core) Core
	Delete(todoCore Core) Core
	FindAll() []Core
	FindById(todoCore Core) Core
}
