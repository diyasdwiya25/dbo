package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type Customers interface {
	Show() ([]entity.Customers, error)
	ShowByLimit(limit int, offset int) ([]entity.Customers, error)
	Store(customer entity.Customers) (entity.Customers, error)
	Update(customer entity.Customers) (entity.Customers, error)
	FindById(id int) (entity.Customers, error)
	Delete(id int) (bool, error)
}

type customers struct {
	db *sql.DB
}

func NewCustomersRepository(db *sql.DB) *customers {
	return &customers{db}
}

func (r *customers) Show() ([]entity.Customers, error) {
	var customer []entity.Customers
	rows, err := r.db.Query(`SELECT id, email, name, address, city, state, postal_code, country, phone, created_at, updated_at from customers`)

	if err != nil {
		return customer, err
	}

	for rows.Next() {
		var customerRow entity.Customers

		err := rows.Scan(&customerRow.Id, &customerRow.Email, &customerRow.Name, &customerRow.Address, &customerRow.City, &customerRow.State, &customerRow.PostalCode, &customerRow.Country, &customerRow.Phone, &customerRow.CreatedAt, &customerRow.UpdatedAt)

		if err != nil {
			return customer, err
		}

		customer = append(customer, customerRow)
	}
	return customer, nil
}

func (r *customers) ShowByLimit(limit int, offset int) ([]entity.Customers, error) {
	var customer []entity.Customers
	rows, err := r.db.Query(`SELECT id, email, name, address, city, state, postal_code, country, phone, created_at, updated_at from customers limit ? offset ?`, limit, offset)

	if err != nil {
		return customer, err
	}

	for rows.Next() {
		var customerRow entity.Customers

		err := rows.Scan(&customerRow.Id, &customerRow.Email, &customerRow.Name, &customerRow.Address, &customerRow.City, &customerRow.State, &customerRow.PostalCode, &customerRow.Country, &customerRow.Phone, &customerRow.CreatedAt, &customerRow.UpdatedAt)

		if err != nil {
			return customer, err
		}

		customer = append(customer, customerRow)
	}
	return customer, nil
}

func (r *customers) Store(customer entity.Customers) (entity.Customers, error) {
	query := `INSERT INTO customers (email, name, address, city, state, postal_code, country, phone, created_at) 
	VALUES (?,?,?,?,?,?,?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &customer.Email, &customer.Name, &customer.Address, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.Phone, &customer.CreatedAt)
	if err != nil {
		return customer, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return customer, err
	}
	customer.Id = int(lastId)
	return customer, nil
}

func (r *customers) Update(customer entity.Customers) (entity.Customers, error) {
	query := `UPDATE customers SET email = ?, name = ?, address = ?, city = ?, state = ?, postal_code = ?, country = ?, phone = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &customer.Email, &customer.Name, &customer.Address, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.Phone, &customer.Id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customers) FindById(id int) (entity.Customers, error) {
	var customer entity.Customers

	err := r.db.QueryRow(`SELECT id, email, name, address, city, state, postal_code, country, phone, created_at, updated_at from customers WHERE id = ?`, id).
		Scan(&customer.Id, &customer.Email, &customer.Name, &customer.Address, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customers) Delete(id int) (bool, error) {
	query := `DELETE FROM customers where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
