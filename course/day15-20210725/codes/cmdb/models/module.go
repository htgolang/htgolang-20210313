package models

import "time"

type Module struct {
	Id        int64
	Key       string
	Title     string
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
