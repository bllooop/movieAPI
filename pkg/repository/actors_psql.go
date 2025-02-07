package repository

import (
	"database/sql"
	"errors"
	"fmt"
	movieapi "movieapi"
	"strings"

	"github.com/lib/pq"
)

type ActorPostgres struct {
	db *sql.DB
}

func NewActorPostgres(db *sql.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (r *ActorPostgres) CreateActor(userRole string, list movieapi.ActorList) (int, error) {
	if userRole == "0" {
		return 0, errors.New("access restricted")
	}
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createQuery := fmt.Sprintf(`INSERT INTO %s (name,gender,date) VALUES ($1,$2,$3) RETURNING id`, actorListTable)
	row := tr.QueryRow(createQuery, list.Name, list.Gender, list.Birthdate)
	if err := row.Scan(&id); err != nil {
		tr.Rollback()
		return 0, err
	}
	return id, tr.Commit()
}

// SELECT at.id,at.name,at.gender,at.date,array_agg(mt.title) FROM actorlist at LEFT JOIN movielist mt ON at.name = ANY(mt.actorname) GROUP BY (at.id, at.name, at.gender,at.date);
func (r *ActorPostgres) ListActors() ([]movieapi.ActorList, error) {
	var lists []movieapi.ActorList
	query := fmt.Sprintf(`SELECT at.id,at.name,at.gender,at.date,array_agg(mt.title)filter (where mt.title is not null) FROM %s at LEFT JOIN %s mt ON at.name = ANY(mt.actorname) GROUP BY (at.id, at.name, at.gender,at.date)`, actorListTable, movieListTable)
	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.ActorList{}
		err := res.Scan(&k.Id, &k.Name, &k.Gender, &k.Birthdate, pq.Array(&k.Movielist))
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

func (r *ActorPostgres) Update(userRole string, actorId int, input movieapi.UpdateActorListInput) error {
	if userRole == "0" {
		return errors.New("access restricted")
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Gender != nil {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *input.Gender)
		argId++
	}
	if input.Birthdate != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Birthdate)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", actorListTable, setQuery, argId)
	args = append(args, actorId)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *ActorPostgres) Delete(userRole string, actorId int) error {
	if userRole == "0" {
		return errors.New("access restricted")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", actorListTable)
	_, err := r.db.Exec(query, actorId)
	return err
}
