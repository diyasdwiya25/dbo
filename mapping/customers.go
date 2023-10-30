package mapping

import (
	"dbo/entity"
	"dbo/helper"
	"time"
)

type customersWithPageFormater struct {
	Customer   []customersFormater `json:"customer"`
	Pagination helper.Pagination   `json:"pagination"`
}

type customersFormater struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postalCode"`
	Country    string    `json:"country"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func CustomersRow(customers entity.Customers) customersFormater {
	formater := customersFormater{
		Id:         customers.Id,
		Name:       customers.Name,
		Email:      customers.Email,
		Address:    customers.Address,
		City:       customers.City,
		State:      customers.State,
		PostalCode: customers.PostalCode,
		Country:    customers.Country,
		Phone:      customers.Phone,
		CreatedAt:  customers.CreatedAt,
		UpdatedAt:  customers.UpdatedAt,
	}

	return formater
}

func Customers(customers []entity.Customers, pagination helper.Pagination) customersWithPageFormater {

	if len(customers) == 0 {
		return customersWithPageFormater{}
	}

	var customersArray []customersFormater
	for _, row := range customers {
		customersRow := CustomersRow(row)
		customersArray = append(customersArray, customersRow)
	}

	formater := customersWithPageFormater{
		Customer:   customersArray,
		Pagination: pagination,
	}

	return formater
}
