package repository

import (
	"database/sql"
	"fmt"
)

const (
	userListTable    = "userlist"
	actorListTable   = "actorlist"
	movieItemTable   = "movieitem"
	actorItemTable   = "actoritem"
	movieListTable   = "movielist"
	movListItemTable = "movlistitem"
)
const (
	host   = "db"
	port   = "5436"
	user   = "postgres"
	dbname = "postgres"
)

func NewPostgresDB(password string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, password, dbname))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
