package repository

import "database/sql"

type Home interface {
}

type Repository struct {
	Home
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
