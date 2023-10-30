package entity

import "time"

type Orders struct {
	Id         int
	Invoice    string
	CustomerId int
	UserId     int
	Total      int
	ShippedAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
