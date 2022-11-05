package activityhandler

type Response struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"deleted_at"`
}
