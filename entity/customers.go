package entity

import "time"

type Customers struct {
	Id         int
	Email      string
	Name       string
	Address    string
	City       string
	State      string
	PostalCode string
	Country    string
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
