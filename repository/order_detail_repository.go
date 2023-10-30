package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type OrderDetail interface {
	Show() ([]entity.OrderDetail, error)
	Store(orderDetail entity.OrderDetail) (entity.OrderDetail, error)
	Update(orderDetail entity.OrderDetail) (entity.OrderDetail, error)
	FindById(id int) (entity.OrderDetail, error)
	Delete(id int) (bool, error)
	DeleteByOrderId(orderId int) (bool, error)
	ShowByOrderId(orderId int) ([]entity.OrderDetail, error)
}

type orderDetail struct {
	db *sql.DB
}

func NewOrderDetailRepository(db *sql.DB) *orderDetail {
	return &orderDetail{db}
}

func (r *orderDetail) Show() ([]entity.OrderDetail, error) {
	var orderDetail []entity.OrderDetail
	rows, err := r.db.Query(`SELECT id, order_id, product_id, qty, price, created_at, updated_at from order_detail`)

	if err != nil {
		return orderDetail, err
	}

	for rows.Next() {
		var orderDetailRow entity.OrderDetail

		err := rows.Scan(&orderDetailRow.Id, &orderDetailRow.OrderId, &orderDetailRow.ProductId, &orderDetailRow.Qty, &orderDetailRow.Price, &orderDetailRow.CreatedAt, &orderDetailRow.UpdatedAt)

		if err != nil {
			return orderDetail, err
		}

		orderDetail = append(orderDetail, orderDetailRow)
	}
	return orderDetail, nil
}

func (r *orderDetail) Store(orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	query := `INSERT INTO order_detail (order_id, product_id, qty, price, created_at) 
	VALUES (?,?,?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.Qty, &orderDetail.Price, &orderDetail.CreatedAt)
	if err != nil {
		return orderDetail, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return orderDetail, err
	}
	orderDetail.Id = int(lastId)
	return orderDetail, nil
}

func (r *orderDetail) Update(orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	query := `UPDATE order_detail SET order_id = ?, product_id = ?, qty = ?, price = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.Qty, &orderDetail.Price, &orderDetail.Id)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (r *orderDetail) FindById(id int) (entity.OrderDetail, error) {
	var orderDetail entity.OrderDetail

	err := r.db.QueryRow(`SELECT id, order_id, product_id, qty, price, created_at, updated_at from order_detail WHERE id = ?`, id).
		Scan(&orderDetail.Id, &orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.Qty, &orderDetail.Price, &orderDetail.CreatedAt, &orderDetail.UpdatedAt)

	if err != nil {
		return orderDetail, err
	}

	return orderDetail, nil
}

func (r *orderDetail) Delete(id int) (bool, error) {
	query := `DELETE FROM order_detail where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *orderDetail) DeleteByOrderId(orderId int) (bool, error) {
	query := `DELETE FROM order_detail where order_id = ?`

	_, err := r.db.ExecContext(context.Background(), query, orderId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *orderDetail) ShowByOrderId(orderId int) ([]entity.OrderDetail, error) {
	var orderDetail []entity.OrderDetail
	rows, err := r.db.Query(`SELECT id, order_id, product_id, qty, price, created_at, updated_at from order_detail where order_id = ?`, orderId)

	if err != nil {
		return orderDetail, err
	}

	for rows.Next() {
		var orderDetailRow entity.OrderDetail

		err := rows.Scan(&orderDetailRow.Id, &orderDetailRow.OrderId, &orderDetailRow.ProductId, &orderDetailRow.Qty, &orderDetailRow.Price, &orderDetailRow.CreatedAt, &orderDetailRow.UpdatedAt)

		if err != nil {
			return orderDetail, err
		}

		orderDetail = append(orderDetail, orderDetailRow)
	}
	return orderDetail, nil
}
