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
	query := fmt.Sprintf(`INSERT INTO %s (username,password,role) VALUES ($1,$2,$3) RETURNING id`, userListTable)
	row := r.db.QueryRow(query, user.UserName, user.Password, 1)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) SignUser(username, password string) (movieapi.User, error) {
	var user movieapi.User
	query := fmt.Sprintf(`SELECT id,username,password,role FROM %s WHERE username=$1 AND password=$2`, userListTable)
	res := r.db.QueryRow(query, username, password)
	err := res.Scan(&user.Id, &user.UserName, &user.Password, &user.Role)
	return user, err
}

/*func (r *AuthPostgres) GetUserRole(username string) (string, error) {
	var role string
	query := fmt.Sprintf(`SELECT role FROM %s WHERE username=$1`, userListTable)
	res := r.db.QueryRow(query, username)
	err := res.Scan(&role)
	return role, err
}*/
