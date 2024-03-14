package repository

import (
	"database/sql"
	"fmt"
	movieapi "movieAPI"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user movieapi.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (username,password,role) VALUES ($1,$2,$3) RETURNING id`, "userlist")
	row := r.db.QueryRow(query, user.UserName, user.Password, 1)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) SignUser(username, password string) (int, error) {
	//var user movieapi.User
	var id int
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password=$2`, "userlist")
	res := r.db.QueryRow(query, username, password)
	err := res.Scan(&id)
	return id, err
}
