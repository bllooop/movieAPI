package repository

import (
	"database/sql"
	"fmt"
	movieapi "movieapi"
	"strings"
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
func (r *MoviePostgres) ListMovies(order string) ([]movieapi.MovieList, error) {
	var lists []movieapi.MovieList
	if order == "" {
		order = "rating"
	}
	query := fmt.Sprintf(`SELECT id,title,rating,date FROM %s ORDER BY %s DESC`, movieListTable, order)
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
func (r *MoviePostgres) Update(role, movId int, input movieapi.UpdateMovieListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}
	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Date)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", movieListTable, setQuery, argId)
	args = append(args, movId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *MoviePostgres) Delete(role, movId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", movieListTable)
	_, err := r.db.Exec(query, movId)
	return err
}
