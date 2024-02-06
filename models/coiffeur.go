package models

type coiffeur struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Salon    string `json:"salon"`
}

type coiffeurStoreInterface interface {
	GetCoiffeur() ([]Salon, error)
	AddCoiffeur(item Salon) (int, error)
	DeleteCoiffeur(id int) error
}
