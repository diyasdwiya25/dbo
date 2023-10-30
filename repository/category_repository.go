package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type Category interface {
	Show() ([]entity.Category, error)
	Store(category entity.Category) (entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	FindById(id int) (entity.Category, error)
	Delete(id int) (bool, error)
}

type category struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *category {
	return &category{db}
}

func (r *category) Show() ([]entity.Category, error) {
	var category []entity.Category
	rows, err := r.db.Query(`SELECT id, name, description, created_at, updated_at from category`)

	if err != nil {
		return category, err
	}

	for rows.Next() {
		var categoryRow entity.Category

		err := rows.Scan(&categoryRow.Id, &categoryRow.Name, &categoryRow.Description, &categoryRow.CreatedAt, &categoryRow.UpdatedAt)

		if err != nil {
			return category, err
		}

		category = append(category, categoryRow)
	}
	return category, nil
}

func (r *category) Store(category entity.Category) (entity.Category, error) {
	query := `INSERT INTO category (name, description, created_at) 
	VALUES (?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &category.Name, &category.Description, &category.CreatedAt)
	if err != nil {
		return category, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return category, err
	}
	category.Id = int(lastId)
	return category, nil
}

func (r *category) Update(category entity.Category) (entity.Category, error) {
	query := `UPDATE category SET name = ?, description = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &category.Name, &category.Description, &category.Id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *category) FindById(id int) (entity.Category, error) {
	var category entity.Category

	err := r.db.QueryRow(`SELECT id, name, description, created_at, updated_at from category WHERE id = ?`, id).
		Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *category) Delete(id int) (bool, error) {
	query := `DELETE FROM category where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
