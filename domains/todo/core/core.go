package todocore

import "time"

type Core struct {
	Id              uint
	ActivityGroupId uint
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
