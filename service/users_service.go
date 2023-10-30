package service

import (
	"dbo/entity"
	"dbo/payload"
	"dbo/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Users interface {
	Show() ([]entity.Users, error)
	Create(input payload.UsersPayload) (entity.Users, error)
	Update(id int, input payload.UsersPayload) (entity.Users, error)
	Find(id int) (entity.Users, error)
	Delete(id int) (bool, error)
	Login(input payload.Login) (entity.Users, error)
}

type users struct {
	repository repository.Users
}

func NewUsersService(repository repository.Users) *users {
	return &users{repository}
}

func (s *users) Show() ([]entity.Users, error) {
	users, err := s.repository.Show()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *users) Create(input payload.UsersPayload) (entity.Users, error) {
	now := time.Now()
	users := entity.Users{}

	password := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return users, err
	}

	users.Name = input.Name
	users.Email = input.Email
	users.Password = string(hashedPassword)
	users.CreatedAt = now

	users, err = s.repository.Store(users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *users) Update(id int, input payload.UsersPayload) (entity.Users, error) {
	users, err := s.repository.FindById(id)
	if err != nil {
		return users, err
	}

	password := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return users, err
	}

	users.Name = input.Name
	users.Email = input.Email
	users.Password = string(hashedPassword)

	users, err = s.repository.Update(users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *users) Find(id int) (entity.Users, error) {
	users, err := s.repository.FindById(id)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *users) Delete(id int) (bool, error) {
	users, err := s.repository.Delete(id)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *users) Login(input payload.Login) (entity.Users, error) {
	email := input.Email
	password := input.Password
	users, err := s.repository.FindByEmail(email)
	if err != nil {
		return users, err
	}

	if users.Id == 0 {
		return users, errors.New("email dan password tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return users, err
	}
	return users, nil
}
