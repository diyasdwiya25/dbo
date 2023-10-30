package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"time"
)

type Category interface {
	Show() ([]entity.Category, error)
	Create(input payload.CategoryPayload) (entity.Category, error)
	Update(id int, input payload.CategoryPayload) (entity.Category, error)
	Find(id int) (entity.Category, error)
	Delete(id int) (bool, error)
}

type category struct {
	repository repository.Category
}

func NewCategoryService(repository repository.Category) *category {
	return &category{repository}
}

func (s *category) Show() ([]entity.Category, error) {
	category, err := s.repository.Show()
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *category) Create(input payload.CategoryPayload) (entity.Category, error) {
	now := time.Now()
	category := entity.Category{}

	category.Name = input.Name
	category.Description = input.Description
	category.CreatedAt = now

	category, err := s.repository.Store(category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *category) Update(id int, input payload.CategoryPayload) (entity.Category, error) {
	category, err := s.repository.FindById(id)
	if err != nil {
		return category, err
	}

	category.Name = input.Name
	category.Description = input.Description

	category, err = s.repository.Update(category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *category) Find(id int) (entity.Category, error) {
	category, err := s.repository.FindById(id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *category) Delete(id int) (bool, error) {
	category, err := s.repository.Delete(id)
	if err != nil {
		return category, err
	}
	return category, nil
}
