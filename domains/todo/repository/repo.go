package todorepo

import (
	todocore "devcode/domains/todo/core"
	todomodel "devcode/domains/todo/model"

	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *todoRepo {
	return &todoRepo{
		db: db,
	}
}

func (r *todoRepo) Insert(todoCore todocore.Core) (todocore.Core, error) {
	model := todomodel.Todo{
		Title:           todoCore.Title,
		ActivityGroupId: todoCore.ActivityGroupId,
		Priority:        "very-high",
		IsActive:        true,
	}

	tx := r.db.Model(&todomodel.Todo{}).Create(&model)
	if tx.Error != nil {
		return todocore.Core{}, tx.Error
	}

	return todomodel.ToCore(model), nil
}

func (r *todoRepo) Update(todoCore todocore.Core) (todocore.Core, error) {
	model := todomodel.Todo{
		Title:           todoCore.Title,
		ActivityGroupId: todoCore.ActivityGroupId,
		Priority:        "very-high",
		IsActive:        todoCore.IsActive,
	}

	tx := r.db.Model(&todomodel.Todo{}).Where("id", todoCore.Id)

	if todoCore.Title != "" {
		tx.Select("title", "is_active")
	} else {
		tx.Select("is_active")
	}

	tx.Updates(&model)
	if tx.Error != nil {
		return todocore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return todocore.Core{}, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(&todomodel.Todo{}).First(&model, todoCore.Id)
	if tx.Error != nil {
		return todocore.Core{}, tx.Error
	}

	return todomodel.ToCore(model), nil
}

func (r *todoRepo) Delete(todoCore todocore.Core) (todocore.Core, error) {
	model := todomodel.Todo{
		Title:           todoCore.Title,
		ActivityGroupId: todoCore.ActivityGroupId,
		Priority:        "very-high",
		IsActive:        todoCore.IsActive,
	}

	tx := r.db.Model(&todomodel.Todo{}).Where("id", todoCore.Id).Delete(&model)
	if tx.Error != nil {
		return todocore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return todocore.Core{}, gorm.ErrRecordNotFound
	}

	return todomodel.ToCore(model), nil
}

func (r *todoRepo) GetAll(todoCore todocore.Core) ([]todocore.Core, error) {
	modelList := []todomodel.Todo{}

	tx := r.db.Model(&todomodel.Todo{}).Unscoped()

	if todoCore.ActivityGroupId > 0 {
		tx.Where("activity_group_id", todoCore.ActivityGroupId)
	}

	tx.Find(&modelList)
	if tx.Error != nil {
		return []todocore.Core{}, tx.Error
	}

	coreList := []todocore.Core{}
	for _, data := range modelList {
		coreList = append(coreList, todomodel.ToCore(data))
	}

	return coreList, nil
}

func (r *todoRepo) GetById(todoCore todocore.Core) (todocore.Core, error) {
	model := todomodel.Todo{}
	tx := r.db.Model(&todomodel.Todo{}).First(&model, todoCore.Id)
	if tx.Error != nil {
		return todocore.Core{}, tx.Error
	}

	return todomodel.ToCore(model), nil
}
