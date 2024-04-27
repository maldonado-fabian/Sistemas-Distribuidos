package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Rut      string `json:"rut"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
