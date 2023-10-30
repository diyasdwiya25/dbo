package entity

import "time"

type Users struct {
	Id            int
	Email         string
	Name          string
	Password      string
	RememberToken string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
