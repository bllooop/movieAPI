package repository

import (
	"database/sql"
)

func NewPostgresDB(database string) (*sql.DB, error) {
	db, err := sql.Open("postgres", database)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
