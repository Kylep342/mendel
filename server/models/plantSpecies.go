package models

import (
	"time"
)

type PlantSpecies struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Taxon     string    `db:"taxon"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
