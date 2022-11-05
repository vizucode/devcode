package activityhandler

type Request struct {
	Email string `json:"email" form:"email"`
	Title string `json:"title" form:"title" validate:"required"`
}
