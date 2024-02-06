package models

import "embed"

var EmbedTemplates embed.FS

type Salon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Tel     int    `json:"tel"`
	Adresse string `json:"adresse"`
}

type SalonStoreInterface interface {
	GetSalons() ([]Salon, error)
	AddSalon(item Salon) (int, error)
	DeleteSalon(id int) error
	ToggleSalon(id int) error
}
