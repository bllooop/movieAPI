package movieapi

type User struct {
	Id       int    `json:"-"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
