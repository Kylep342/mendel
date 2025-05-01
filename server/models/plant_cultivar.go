package models

import (
	"time"
)

type PlantCultivar struct {
	Id        string    `db:"id" json:"id"`
	SpeciesId string    `db:"species_id" json:"species_id"`
	Cultivar  string    `db:"cultivar" json:"cultivar"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Genetics  struct{}  `db:"genetics" json:"genetics"`
}
