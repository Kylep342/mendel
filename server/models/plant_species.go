package models

import (
	"time"
)

type PlantSpecies struct {
	Id        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Taxon     string    `db:"taxon" json:"taxon"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
