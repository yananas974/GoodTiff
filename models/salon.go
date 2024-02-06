package models

import "embed"

var EmbedTemplates embed.FS

type Salon struct {
	ID      int    `json:"id"`
	Name    string `json:"Name"`
	Tel     int    `json:"Tel"`
	Adresse string `json:"Adresse"`
}

type SalonStoreInterface interface {
	GetSalons() ([]Salon, error)
	AddSalon(item Salon) (int, error)
	DeleteSalon(id int) error
	ToggleSalon(id int) error
}
