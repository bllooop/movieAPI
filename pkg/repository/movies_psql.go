package repository

import (
	"database/sql"
	"fmt"
	movieapi "movieAPI"
)

type MoviePostgres struct {
	db *sql.DB
}

func NewMoviePostgres(db *sql.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}
func (r *MoviePostgres) ListMovies() (movieapi.User, error) {
	var user movieapi.User
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username = $1`, "userlist")
	res := r.db.QueryRow(query, "aaaa")
	err := res.Scan(&user)
	return user, err
}
