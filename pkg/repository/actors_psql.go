package repository

import (
	"database/sql"
	"fmt"
	movieapi "movieapi"
	"strings"
)

type ActorPostgres struct {
	db *sql.DB
}

func NewActorPostgres(db *sql.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (r *ActorPostgres) CreateActor(userRole string, list movieapi.ActorList) (int, error) {
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
func (r *ActorPostgres) ListActors() ([]movieapi.ActorList, error) {
	var lists []movieapi.ActorList
	query := fmt.Sprintf(`SELECT id,name,gender,date FROM %s`, actorListTable)
	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		k := movieapi.ActorList{}
		err := res.Scan(&k.Id, &k.Name, &k.Gender, &k.Birthdate)
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
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", actorListTable)
	_, err := r.db.Exec(query, actorId)
	return err
}
