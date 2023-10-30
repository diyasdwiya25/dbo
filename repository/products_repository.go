package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type Products interface {
	Show() ([]entity.Products, error)
	Store(products entity.Products) (entity.Products, error)
	Update(products entity.Products) (entity.Products, error)
	FindById(id int) (entity.Products, error)
	Delete(id int) (bool, error)
}

type products struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *products {
	return &products{db}
}

func (r *products) Show() ([]entity.Products, error) {
	var products []entity.Products
	rows, err := r.db.Query(`SELECT id, name, description, stock, price, category_id, created_at, updated_at from products`)

	if err != nil {
		return products, err
	}

	for rows.Next() {
		var productsRow entity.Products

		err := rows.Scan(&productsRow.Id, &productsRow.Name, &productsRow.Description, &productsRow.Stock, &productsRow.Price, &productsRow.CategoryId, &productsRow.CreatedAt, &productsRow.UpdatedAt)

		if err != nil {
			return products, err
		}

		products = append(products, productsRow)
	}
	return products, nil
}

func (r *products) Store(products entity.Products) (entity.Products, error) {
	query := `INSERT INTO products (name, description, stock, price, category_id, created_at) 
	VALUES (?,?,?,?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &products.Name, &products.Description, &products.Stock, &products.Price, &products.CategoryId, &products.CreatedAt)
	if err != nil {
		return products, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return products, err
	}
	products.Id = int(lastId)
	return products, nil
}

func (r *products) Update(products entity.Products) (entity.Products, error) {
	query := `UPDATE products SET name = ?, description = ?, stock = ?, price = ?, category_id = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &products.Name, &products.Description, &products.Stock, &products.Price, &products.CategoryId, &products.Id)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *products) FindById(id int) (entity.Products, error) {
	var products entity.Products

	err := r.db.QueryRow(`SELECT id, name, description, stock, price, category_id, created_at, updated_at from products WHERE id = ?`, id).
		Scan(&products.Id, &products.Name, &products.Description, &products.Stock, &products.Price, &products.CategoryId, &products.CreatedAt, &products.UpdatedAt)

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *products) Delete(id int) (bool, error) {
	query := `DELETE FROM products where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
