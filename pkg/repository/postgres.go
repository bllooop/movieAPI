package repository

import (
	"database/sql"
)

const (
	userListTable    = "userlist"
	actorListTable   = "actorlist"
	movieItemTable   = "movieitem"
	actorItemTable   = "actoritem"
	movieListTable   = "movielist"
	movListItemTable = "movlistitem"
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
