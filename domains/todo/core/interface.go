package todocore

type IRepoTodo interface {
	Create(todoCore Core) (Core, error)
	Update(todoCore Core) (Core, error)
	Delete(todoCore Core) (Core, error)
	GetAll() ([]Core, error)
	GetById(todoCore Core) (Core, error)
}

type IServiceTodo interface {
	Insert(todoCore Core) Core
	Update(todoCore Core) Core
	Delete(todoCore Core) Core
	FindAll() []Core
	FindById(todoCore Core) Core
}
