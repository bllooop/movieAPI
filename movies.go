package movieapi

import "errors"

type MovieList struct {
	Id     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Rating int    `json:"rating"`
	Date   string `json:"date"`
}

type MovieItem struct {
	Id          int      `json:"id" db:"id"`
	Title       string   `json:"title" db:"title" binding:"required"`
	Description string   `json:"description" db:"description"`
	Rating      int      `json:"rating" db:"rating"`
	Date        string   `json:"date" db:"date"`
	ActorName   []string `json:"actorname" db:"actorname"`
}
type MovListsItem struct {
	Id          int
	MovListId   int
	MovItemId   int
	ActorListId int
	ActorItemId int
}
type ActorItem struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Gender    string `json:"gender" db:"gender"`
	Birthdate string `json:"date" db:"date"`
}

type ActorList struct {
	Id        int    `json:"id"`
	Name      string `json:"name"  binding:"required"`
	Gender    string `json:"gender" db:"gender"`
	Birthdate string `json:"date" db:"date"`
}

type UpdateMovieListInput struct {
	Title  *string `json:"title"`
	Rating *string `json:"rating"`
	Date   *string `json:"date"`
}

func (i UpdateMovieListInput) Validation() error {
	if i.Title == nil && i.Rating == nil && i.Date == nil {
		return errors.New("update params have no values")
	}
	return nil
}

type UpdateMovieItemInput struct {
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Rating      *int      `json:"rating"`
	Date        *string   `json:"date"`
	ActorName   *[]string `json:"actorname"`
}

func (i UpdateMovieItemInput) Validation() error {
	if i.Title == nil && i.Description == nil && i.Rating == nil && i.Date == nil && i.ActorName == nil {
		return errors.New("update params have no values")
	}
	return nil
}

type UpdateActorListInput struct {
	Name      *string `json:"name"`
	Gender    *string `json:"gender"`
	Birthdate *string `json:"date"`
}

func (i UpdateActorListInput) ValidationAct() error {
	if i.Name == nil && i.Gender == nil && i.Birthdate == nil {
		return errors.New("update params have no values")
	}
	return nil
}
