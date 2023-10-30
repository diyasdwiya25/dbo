package entity

import "time"

type OrderDetail struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
