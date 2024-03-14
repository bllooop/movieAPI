package movieapi

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
