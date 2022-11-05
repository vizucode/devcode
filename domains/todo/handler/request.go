package todohandler

type Request struct {
	Title           string `json:"title" form:"title" validate:"required"`
	ActivityGroupId uint   `json:"activity_group_id" form:"activity_group_id" validate:"required"`
}

type RequestUpdate struct {
	Title    string `json:"title" form:"title"`
	IsActive bool   `json:"is_active" form:"is_active"`
}
