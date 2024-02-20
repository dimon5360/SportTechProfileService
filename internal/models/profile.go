package models

import "time"

type Profile struct {
	Id         uint64
	Username   string
	Firstname  string
	Lastname   string
	UserId     uint64
	Created_at time.Time
	Updated_at time.Time
}
