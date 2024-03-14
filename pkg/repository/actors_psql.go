package repository

import (
	"database/sql"
	"fmt"
	movieapi "movieAPI"
)

type ActorPostgres struct {
	db *sql.DB
}

func NewActorPostgres(db *sql.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (r *ActorPostgres) CreateActor(role int, list movieapi.ActorList) (int, error) {
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
	query := fmt.Sprintf(`SELECT id,title,rating,date FROM %s ORDER BY rating DESC`, actorListTable)
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
