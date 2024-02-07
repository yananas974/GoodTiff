package models

type Coiffeur struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Salon    string `json:"salon"`
}

type CoiffeurStoreInterface interface {
	GetCoiffeur() ([]Coiffeur, error)
	AddCoiffeur(item Coiffeur) (int, error)
	DeleteCoiffeur(id int) error
}
