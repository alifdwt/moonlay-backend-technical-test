package models

import "time"

type Migration struct {
	Id int
	Name string
	CreatedAt time.Time
}