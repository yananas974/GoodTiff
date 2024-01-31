package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Tel      int    `json:"Tel"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
