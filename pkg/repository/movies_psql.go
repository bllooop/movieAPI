package repository

import (
	"database/sql"
	"errors"
	"fmt"
	movieapi "movieapi"
	"strings"

	"github.com/lib/pq"
)

type MoviePostgres struct {
	db *sql.DB
}

func NewMoviePostgres(db *sql.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (r *MoviePostgres) Create(userRole string, list movieapi.MovieList) (int, error) {
	if userRole == "0" {
		return 0, errors.New("access restricted")
	}
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createQuery := fmt.Sprintf(`INSERT INTO %s (title,rating,date,description,actorname) VALUES ($1,$2,$3,$4,$5) RETURNING id`, movieListTable)
	row := tr.QueryRow(createQuery, list.Title, list.Rating, list.Date, list.Description, pq.Array(list.ActorName))
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
	query := fmt.Sprintf(`SELECT id,title,rating,date,description,actorname FROM %s ORDER BY %s DESC`, movieListTable, order)
	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.MovieList{}
		err := res.Scan(&k.Id, &k.Title, &k.Rating, &k.Date, &k.Description, pq.Array(&k.ActorName))
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
	res, err := r.db.Query("SELECT id,title,rating,date,description,actorname FROM movielist WHERE title LIKE '%' || $1 || '%' OR ARRAY_TO_STRING(actorname, ',') LIKE '%' || $1 || '%'", movieName)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.MovieList{}
		err := res.Scan(&k.Id, &k.Title, &k.Rating, &k.Date, &k.Description, pq.Array(&k.ActorName))
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
func (r *MoviePostgres) Update(userRole string, movId int, input movieapi.UpdateMovieListInput) error {
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

func (r *MoviePostgres) Delete(userRole string, movId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", movieListTable)
	_, err := r.db.Exec(query, movId)
	return err
}
