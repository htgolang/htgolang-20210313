package models

import "time"

type Module struct {
	Id        int64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	Path       string
	Controller string
	Action     string
	Method     string

	IsMenu bool
}
