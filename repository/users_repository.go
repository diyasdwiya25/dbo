package repository

import (
	"context"
	"database/sql"
	"dbo/entity"
)

type Users interface {
	Show() ([]entity.Users, error)
	Store(users entity.Users) (entity.Users, error)
	Update(users entity.Users) (entity.Users, error)
	FindById(id int) (entity.Users, error)
	Delete(id int) (bool, error)
	FindByEmail(email string) (entity.Users, error)
}

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (r *users) Show() ([]entity.Users, error) {
	var users []entity.Users
	rows, err := r.db.Query(`SELECT id, email, name, password, created_at, updated_at from users`)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var usersRow entity.Users

		err := rows.Scan(&usersRow.Id, &usersRow.Email, &usersRow.Name, &usersRow.Password, &usersRow.CreatedAt, &usersRow.UpdatedAt)

		if err != nil {
			return users, err
		}

		users = append(users, usersRow)
	}
	return users, nil
}

func (r *users) Store(users entity.Users) (entity.Users, error) {
	query := `INSERT INTO users (email, name, password, created_at) 
	VALUES (?,?,?,?)`

	res, err := r.db.ExecContext(context.Background(), query, &users.Email, &users.Name, &users.Password, &users.CreatedAt)
	if err != nil {
		return users, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return users, err
	}
	users.Id = int(lastId)
	return users, nil
}

func (r *users) Update(users entity.Users) (entity.Users, error) {
	query := `UPDATE users SET email = ?, name = ?, password = ?
	WHERE id = ?`

	_, err := r.db.ExecContext(context.Background(), query, &users.Email, &users.Name, &users.Password, &users.Id)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *users) FindById(id int) (entity.Users, error) {
	var users entity.Users

	err := r.db.QueryRow(`SELECT id, email, name, password from users WHERE id = ?`, id).
		Scan(&users.Id, &users.Email, &users.Name, &users.Password)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *users) FindByEmail(email string) (entity.Users, error) {
	var users entity.Users

	err := r.db.QueryRow(`SELECT id, email, name, password from users WHERE email = ?`, email).
		Scan(&users.Id, &users.Email, &users.Name, &users.Password)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *users) Delete(id int) (bool, error) {
	query := `DELETE FROM users where id = ?`

	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
