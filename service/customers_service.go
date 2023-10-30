package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"time"
)

type Customers interface {
	Show() ([]entity.Customers, error)
	ShowByLimit(limit int, offset int) ([]entity.Customers, error)
	Create(input payload.CustomersPayload) (entity.Customers, error)
	Update(id int, input payload.CustomersPayload) (entity.Customers, error)
	Find(id int) (entity.Customers, error)
	Delete(id int) (bool, error)
}

type customers struct {
	repository repository.Customers
}

func NewCustomersService(repository repository.Customers) *customers {
	return &customers{repository}
}

func (s *customers) Show() ([]entity.Customers, error) {
	customers, err := s.repository.Show()
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customers) ShowByLimit(limit int, offset int) ([]entity.Customers, error) {
	customers, err := s.repository.ShowByLimit(limit, offset)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customers) Create(input payload.CustomersPayload) (entity.Customers, error) {
	now := time.Now()
	customers := entity.Customers{}

	customers.Name = input.Name
	customers.Email = input.Email
	customers.Address = input.Address
	customers.City = input.City
	customers.State = input.State
	customers.PostalCode = input.PostalCode
	customers.Country = input.Country
	customers.Phone = input.Phone
	customers.CreatedAt = now

	customers, err := s.repository.Store(customers)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customers) Update(id int, input payload.CustomersPayload) (entity.Customers, error) {
	customers, err := s.repository.FindById(id)
	if err != nil {
		return customers, err
	}

	customers.Name = input.Name
	customers.Email = input.Email
	customers.Address = input.Address
	customers.City = input.City
	customers.State = input.State
	customers.PostalCode = input.PostalCode
	customers.Country = input.Country
	customers.Phone = input.Phone

	customers, err = s.repository.Update(customers)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customers) Find(id int) (entity.Customers, error) {
	customers, err := s.repository.FindById(id)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customers) Delete(id int) (bool, error) {
	customers, err := s.repository.Delete(id)
	if err != nil {
		return customers, err
	}
	return customers, nil
}
