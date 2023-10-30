package mapping

import (
	"dbo/entity"
	"dbo/helper"
	"time"
)

type ordersWithPageFormater struct {
	Orders     []ordersFormater  `json:"orders"`
	Pagination helper.Pagination `json:"pagination"`
}

type ordersFormater struct {
	Id         int       `json:"id"`
	Invoice    string    `json:"invoice"`
	CustomerId int       `json:"customer"`
	UserId     int       `json:"userId"`
	Total      int       `json:"total"`
	ShippedAt  time.Time `json:"shippedAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func OrdersRow(orders entity.Orders) ordersFormater {
	formater := ordersFormater{
		Id:         orders.Id,
		Invoice:    orders.Invoice,
		CustomerId: orders.CustomerId,
		UserId:     orders.UserId,
		Total:      orders.Total,
		ShippedAt:  orders.ShippedAt,
		CreatedAt:  orders.CreatedAt,
		UpdatedAt:  orders.UpdatedAt,
	}

	return formater
}

func Orders(orders []entity.Orders, pagination helper.Pagination) ordersWithPageFormater {

	if len(orders) == 0 {
		return ordersWithPageFormater{}
	}

	var ordersArray []ordersFormater
	for _, row := range orders {
		ordersRow := OrdersRow(row)
		ordersArray = append(ordersArray, ordersRow)
	}

	formater := ordersWithPageFormater{
		Orders:     ordersArray,
		Pagination: pagination,
	}

	return formater
}

type ordersDetailFormater struct {
	Id          int                  `json:"id"`
	Invoice     string               `json:"invoice"`
	CustomerId  int                  `json:"customer"`
	UserId      int                  `json:"userId"`
	Total       int                  `json:"total"`
	OrderDetail []orderDetailProduct `json:"order_detail"`
	ShippedAt   time.Time            `json:"shippedAt"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

type orderDetailProduct struct {
	ProductId int `json:"productId"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
}

func OrdersDetails(orders entity.Orders, orderDetail []entity.OrderDetail) ordersDetailFormater {

	orderProduct := OrdersDetail(orderDetail)
	formater := ordersDetailFormater{
		Id:          orders.Id,
		Invoice:     orders.Invoice,
		CustomerId:  orders.CustomerId,
		UserId:      orders.UserId,
		Total:       orders.Total,
		OrderDetail: orderProduct,
		ShippedAt:   orders.ShippedAt,
		CreatedAt:   orders.CreatedAt,
		UpdatedAt:   orders.UpdatedAt,
	}

	return formater
}

func OrderDetailRow(ordersDetail entity.OrderDetail) orderDetailProduct {
	formater := orderDetailProduct{
		ProductId: ordersDetail.ProductId,
		Qty:       ordersDetail.Qty,
		Price:     ordersDetail.Price,
	}

	return formater
}

func OrdersDetail(orders []entity.OrderDetail) []orderDetailProduct {

	if len(orders) == 0 {
		return []orderDetailProduct{}
	}

	var ordersArray []orderDetailProduct
	for _, row := range orders {
		ordersRow := OrderDetailRow(row)
		ordersArray = append(ordersArray, ordersRow)
	}

	return ordersArray
}
