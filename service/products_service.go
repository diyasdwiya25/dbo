package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"time"
)

type Products interface {
	Show() ([]entity.Products, error)
	Create(input payload.ProductsPayload) (entity.Products, error)
	Update(id int, input payload.ProductsPayload) (entity.Products, error)
	Find(id int) (entity.Products, error)
	Delete(id int) (bool, error)
}

type products struct {
	repository repository.Products
}

func NewProductsService(repository repository.Products) *products {
	return &products{repository}
}

func (s *products) Show() ([]entity.Products, error) {
	products, err := s.repository.Show()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *products) Create(input payload.ProductsPayload) (entity.Products, error) {
	now := time.Now()
	products := entity.Products{}

	products.Name = input.Name
	products.Description = input.Description
	products.Stock = input.Stock
	products.Price = input.Price
	products.CategoryId = input.CategoryId
	products.CreatedAt = now

	products, err := s.repository.Store(products)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *products) Update(id int, input payload.ProductsPayload) (entity.Products, error) {
	products, err := s.repository.FindById(id)
	if err != nil {
		return products, err
	}

	products.Name = input.Name
	products.Description = input.Description
	products.Stock = input.Stock
	products.Price = input.Price
	products.CategoryId = input.CategoryId

	products, err = s.repository.Update(products)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *products) Find(id int) (entity.Products, error) {
	products, err := s.repository.FindById(id)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *products) Delete(id int) (bool, error) {
	products, err := s.repository.Delete(id)
	if err != nil {
		return products, err
	}
	return products, nil
}
