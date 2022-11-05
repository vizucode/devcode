package activitycore

import "time"

type Core struct {
	Id        uint
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
