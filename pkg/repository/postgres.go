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

/*
const (

	host = "db"   //comment when starting on local without docker-compose
	port = "5432" //comment when starting on local without docker-compose
	//host     = "localhost" //uncomment when starting on local without docker-compose
	//port     = "5432"      //uncomment when starting on local without docker-compose
	user     = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
	password = "54321"

)
*/
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
