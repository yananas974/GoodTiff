package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Tel      int    `json:"Tel"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type UserStoreInterface interface {
	GetUsers() ([]User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
}
