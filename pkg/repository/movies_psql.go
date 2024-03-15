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

func (r *MoviePostgres) Create(role int, list movieapi.MovieList) (int, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createQuery := fmt.Sprintf(`INSERT INTO %s (title,rating,date) VALUES ($1,$2,$3) RETURNING id`, movieListTable)
	row := tr.QueryRow(createQuery, list.Title, list.Rating, list.Date)
	if err := row.Scan(&id); err != nil {
		tr.Rollback()
		return 0, err
	}
	return id, tr.Commit()
}
func (r *MoviePostgres) ListMovies() ([]movieapi.MovieList, error) {
	var lists []movieapi.MovieList
	query := fmt.Sprintf(`SELECT id,title,rating,date FROM %s ORDER BY rating DESC`, movieListTable)
	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.MovieList{}
		err := res.Scan(&k.Id, &k.Title, &k.Rating, &k.Date)
		if err != nil {
			return nil, err
		}
		lists = append(lists, k)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *MoviePostgres) GetByName(movieName string) ([]movieapi.MovieList, error) {
	var list []movieapi.MovieList
	res, err := r.db.Query("SELECT id,title,rating,date FROM movielist WHERE title LIKE '%' || $1 || '%'", movieName)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.MovieList{}
		err := res.Scan(&k.Id, &k.Title, &k.Rating, &k.Date)
		if err != nil {
			return nil, err
		}
		list = append(list, k)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
