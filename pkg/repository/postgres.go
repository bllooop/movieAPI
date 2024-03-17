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
	//host     = "db"
	//port     = "5432"
	host     = "localhost"
	port     = "5433"
	user     = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
	password = "54321"
)

func NewPostgresDB() (*sql.DB, error) {
	//db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, "54321", dbname, sslmode))
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
