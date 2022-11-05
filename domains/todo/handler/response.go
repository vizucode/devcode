package todohandler

import (
	"devcode/config"
	todocore "devcode/domains/todo/core"
)

type Response struct {
	Id              uint   `json:"id"`
	Title           string `json:"title"`
	ActivityGroupId uint   `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

func ToResponse(todoCore todocore.Core) Response {
	return Response{
		Id:              todoCore.Id,
		Title:           todoCore.Title,
		ActivityGroupId: todoCore.ActivityGroupId,
		IsActive:        todoCore.IsActive,
		Priority:        todoCore.Priority,
		CreatedAt:       todoCore.CreatedAt.Format(config.LAYOUT_TIME),
		UpdatedAt:       todoCore.UpdatedAt.Format(config.LAYOUT_TIME),
		DeletedAt:       todoCore.DeletedAt.Format(config.LAYOUT_TIME),
	}
}
