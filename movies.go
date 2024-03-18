package movieapi

import "errors"

type MovieList struct {
	Id          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Rating      int      `json:"rating"`
	Date        string   `json:"date"`
	Description string   `json:"description" db:"description"`
	ActorName   []string `json:"actorname" db:"actorname"`
}

type MovListsItem struct {
	Id          int
	MovListId   int
	ActorListId int
}

type ActorList struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"  binding:"required"`
	Gender    string   `json:"gender" db:"gender"`
	Birthdate string   `json:"date" db:"date"`
	Movielist []string `json:"movies"`
}
type UpdateMovieListInput struct {
	Title       *string   `json:"title"`
	Rating      *int      `json:"rating"`
	Description *string   `json:"description"`
	Date        *string   `json:"date"`
	ActorName   *[]string `json:"actorname"`
}

func (i UpdateMovieListInput) Validation() error {
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
