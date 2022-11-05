package todomodel

import (
	todocore "devcode/domains/todo/core"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title           string
	ActivityGroupId uint
	IsActive        bool
	Priority        string
}

func ToCore(model Todo) todocore.Core {
	return todocore.Core{
		Id:              model.ID,
		ActivityGroupId: model.ActivityGroupId,
		Title:           model.Title,
		IsActive:        model.IsActive,
		Priority:        model.Priority,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		DeletedAt:       model.DeletedAt.Time,
	}
}
