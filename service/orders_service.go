package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"math/rand"
	"strconv"
	"time"
)

type Orders interface {
	Show() ([]entity.Orders, error)
	ShowByLimit(limit int, offset int) ([]entity.Orders, error)
	Create(input payload.OrdersPayload, userId int) (entity.Orders, error)
	Update(id int, input payload.OrdersPayload) (entity.Orders, error)
	Find(id int) (entity.Orders, error)
	Delete(id int) (bool, error)
}

type orders struct {
	repository         repository.Orders
	repositoryProducts repository.Products
}

func NewOrdersService(repository repository.Orders, repositoryProducts repository.Products) *orders {
	return &orders{repository, repositoryProducts}
}

func (s *orders) Show() ([]entity.Orders, error) {
	orders, err := s.repository.Show()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orders) ShowByLimit(limit int, offset int) ([]entity.Orders, error) {
	orders, err := s.repository.ShowByLimit(limit, offset)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orders) Create(input payload.OrdersPayload, userId int) (entity.Orders, error) {
	now := time.Now()
	orders := entity.Orders{}

	invoiceNumber := 100000 + rand.Intn(999999-100000)

	orders.Invoice = "invoice-" + strconv.Itoa(invoiceNumber)
	orders.CustomerId = input.CustomerId
	orders.UserId = userId

	inputDetailOrder := input.Order

	total := 0
	for _, row := range inputDetailOrder {

		productsFindById, err := s.repositoryProducts.FindById(row.ProductId)
		if err != nil {
			return orders, err
		}

		productPrice := productsFindById.Price * row.Qty
		total += productPrice
	}

	orders.Total = total

	dateShippeddParse, err := time.Parse("02-01-2006", input.ShippedAt)
	if err != nil {
		return orders, err
	}

	orders.ShippedAt = dateShippeddParse
	orders.CreatedAt = now

	orders, err = s.repository.Store(orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orders) Update(id int, input payload.OrdersPayload) (entity.Orders, error) {
	orders, err := s.repository.FindById(id)
	if err != nil {
		return orders, err
	}

	total := 0
	inputDetailOrder := input.Order
	for _, row := range inputDetailOrder {

		productsFindById, err := s.repositoryProducts.FindById(row.ProductId)
		if err != nil {
			return orders, err
		}
		productPrice := productsFindById.Price * row.Qty
		total += productPrice
	}

	orders.Total = total

	dateShippeddParse, err := time.Parse("02-01-2006", input.ShippedAt)
	if err != nil {
		return orders, err
	}

	orders.ShippedAt = dateShippeddParse

	orders, err = s.repository.Update(orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orders) Find(id int) (entity.Orders, error) {
	orders, err := s.repository.FindById(id)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orders) Delete(id int) (bool, error) {
	orders, err := s.repository.Delete(id)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
