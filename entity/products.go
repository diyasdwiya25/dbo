package entity

import "time"

type Products struct {
	Id          int
	Name        string
	Description string
	Stock       int
	Price       int
	CategoryId  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
