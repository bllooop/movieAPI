package movieapi

type MovieList struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Rating int    `json:"rating"`
	Date   string `json:"date"`
}

type MovieItem struct {
	Id          int      `json:"id" db:"id"`
	Title       string   `json:"title" db:"title" binding:"required"`
	Description string   `json:"description" db:"description"`
	Rating      int      `json:"rating" db:"price"`
	Date        string   `json:"date" db:"expiration"`
	Actors      []string `json:"actors"`
}

type ActorItem struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"date"`
}
