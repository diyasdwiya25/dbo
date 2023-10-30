package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type Orders interface {
	Show() ([]entity.Orders, error)
	ShowByLimit(limit int, offset int) ([]entity.Orders, error)
	Store(orders entity.Orders) (entity.Orders, error)
	Update(orders entity.Orders) (entity.Orders, error)
	FindById(id int) (entity.Orders, error)
	Delete(id int) (bool, error)
}

type orders struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) *orders {
	return &orders{db}
}

func (r *orders) Show() ([]entity.Orders, error) {
	var orders []entity.Orders
	rows, err := r.db.Query(`SELECT id, invoice, customer_id, user_id, total, shipped_at, created_at, updated_at from orders`)

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var ordersRow entity.Orders

		err := rows.Scan(&ordersRow.Id, &ordersRow.Invoice, &ordersRow.CustomerId, &ordersRow.UserId, &ordersRow.Total, &ordersRow.ShippedAt, &ordersRow.CreatedAt, &ordersRow.UpdatedAt)

		if err != nil {
			return orders, err
		}

		orders = append(orders, ordersRow)
	}
	return orders, nil
}

func (r *orders) ShowByLimit(limit int, offset int) ([]entity.Orders, error) {
	var orders []entity.Orders
	rows, err := r.db.Query(`SELECT id, invoice, customer_id, user_id, total, shipped_at, created_at, updated_at from orders limit ? offset ?`, limit, offset)

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var ordersRow entity.Orders

		err := rows.Scan(&ordersRow.Id, &ordersRow.Invoice, &ordersRow.CustomerId, &ordersRow.UserId, &ordersRow.Total, &ordersRow.ShippedAt, &ordersRow.CreatedAt, &ordersRow.UpdatedAt)

		if err != nil {
			return orders, err
		}

		orders = append(orders, ordersRow)
	}
	return orders, nil
}

func (r *orders) Store(orders entity.Orders) (entity.Orders, error) {
	query := `INSERT INTO orders (invoice, customer_id, user_id, total, shipped_at, created_at) 
	VALUES (?,?,?,?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &orders.Invoice, &orders.CustomerId, &orders.UserId, &orders.Total, &orders.ShippedAt, &orders.CreatedAt)
	if err != nil {
		return orders, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return orders, err
	}
	orders.Id = int(lastId)
	return orders, nil
}

func (r *orders) Update(orders entity.Orders) (entity.Orders, error) {
	query := `UPDATE orders SET invoice = ?, customer_id = ?, user_id = ?, total = ?, shipped_at = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &orders.Invoice, &orders.CustomerId, &orders.UserId, &orders.Total, &orders.ShippedAt, &orders.Id)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *orders) FindById(id int) (entity.Orders, error) {
	var orders entity.Orders

	err := r.db.QueryRow(`SELECT id, invoice, customer_id, user_id, total, shipped_at, created_at, updated_at from orders WHERE id = ?`, id).
		Scan(&orders.Id, &orders.Invoice, &orders.CustomerId, &orders.UserId, &orders.Total, &orders.ShippedAt, &orders.CreatedAt, &orders.UpdatedAt)

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *orders) Delete(id int) (bool, error) {
	query := `DELETE FROM orders where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
