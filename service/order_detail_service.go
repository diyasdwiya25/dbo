package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"time"
)

type OrderDetail interface {
	Show() ([]entity.OrderDetail, error)
	Create(input payload.OrderDetailPayload) (entity.OrderDetail, error)
	Update(id int, input payload.OrderDetailPayload) (entity.OrderDetail, error)
	Find(id int) (entity.OrderDetail, error)
	Delete(id int) (bool, error)
	DeleteByOrderId(orderId int) (bool, error)
	CreateMultiple(input []payload.OrderDetailPayload, orderId int) ([]entity.OrderDetail, error)
	ShowByOrderId(orderId int) ([]entity.OrderDetail, error)
}

type orderDetail struct {
	repository         repository.OrderDetail
	repositoryProducts repository.Products
}

func NewOrderDetailService(repository repository.OrderDetail, repositoryProducts repository.Products) *orderDetail {
	return &orderDetail{repository, repositoryProducts}
}

func (s *orderDetail) Show() ([]entity.OrderDetail, error) {
	orderDetail, err := s.repository.Show()
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) Create(input payload.OrderDetailPayload) (entity.OrderDetail, error) {
	now := time.Now()
	orderDetail := entity.OrderDetail{}

	orderDetail.OrderId = 1
	orderDetail.ProductId = input.ProductId
	orderDetail.Qty = input.Qty
	orderDetail.Price = 0
	orderDetail.CreatedAt = now

	orderDetail, err := s.repository.Store(orderDetail)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) Update(id int, input payload.OrderDetailPayload) (entity.OrderDetail, error) {
	orderDetail, err := s.repository.FindById(id)
	if err != nil {
		return orderDetail, err
	}

	orderDetail.OrderId = 1
	orderDetail.ProductId = input.ProductId
	orderDetail.Qty = input.Qty
	orderDetail.Price = 0

	orderDetail, err = s.repository.Update(orderDetail)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) Find(id int) (entity.OrderDetail, error) {
	orderDetail, err := s.repository.FindById(id)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) Delete(id int) (bool, error) {
	orderDetail, err := s.repository.Delete(id)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) DeleteByOrderId(orderId int) (bool, error) {
	orderDetail, err := s.repository.DeleteByOrderId(orderId)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (s *orderDetail) CreateMultiple(input []payload.OrderDetailPayload, orderId int) ([]entity.OrderDetail, error) {
	now := time.Now()
	orderDetail := entity.OrderDetail{}
	var orderDetailArray []entity.OrderDetail
	for _, row := range input {
		orderDetail.OrderId = orderId
		orderDetail.ProductId = row.ProductId
		orderDetail.Qty = row.Qty

		productsFindById, err := s.repositoryProducts.FindById(row.ProductId)
		if err != nil {
			return orderDetailArray, err
		}

		orderDetail.Price = productsFindById.Price
		orderDetail.CreatedAt = now

		orderDetail, err := s.repository.Store(orderDetail)
		if err != nil {
			return orderDetailArray, err
		}

		orderDetailArray = append(orderDetailArray, orderDetail)
	}
	return orderDetailArray, nil
}

func (s *orderDetail) ShowByOrderId(orderId int) ([]entity.OrderDetail, error) {
	orderDetail, err := s.repository.ShowByOrderId(orderId)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}
