package activitycore

type IRepoActivity interface {
	Insert(activityCore Core) (Core, error)
	Update(activityCore Core) (Core, error)
	Delete(activityCore Core) (Core, error)
	GetAll() ([]Core, error)
	GetSingle(activityCore Core) (Core, error)
}

type IServiceActivity interface {
	Create(activityCore Core) Core
	Update(activityCore Core) Core
	Delete(activityCore Core) Core
	FindAll() []Core
	FindSingle(activityCore Core) Core
}
