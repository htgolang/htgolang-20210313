package models

import "time"

type Alarm struct {
	Id        int64
	AlarmTime *time.Time
	Content   string
	Status    int
}
