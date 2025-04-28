package models

import (
	"time"
)

type PlantCultivar struct {
	Id        string    `db:"id"`
	SpeciesId string    `db:"species_id"`
	Cultivar  string    `db:"cultivar"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Genetics  struct{}  `db:"genetics"`
}
