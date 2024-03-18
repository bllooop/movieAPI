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
	//host = "db"   //comment when starting on local without docker-compose
	//port = "5432" //comment when starting on local without docker-compose
	host     = "localhost" //uncomment when starting on local without docker-compose
	port     = "5433"      //uncomment when starting on local without docker-compose
	user     = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
	password = "54321"
)

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
