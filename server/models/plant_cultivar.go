package models

import (
	"time"
)

type PlantCultivar struct {
	ID        string      `db:"id" json:"id"`
	SpeciesID string      `db:"species_id" json:"species_id"`
	Name      string      `db:"name" json:"name"`
	Cultivar  string      `db:"cultivar" json:"cultivar"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	Genetics  interface{} `db:"genetics" json:"genetics"`
}

func (p *PlantCultivar) GetID() string { return p.ID }

func (p *PlantCultivar) SetID(id string) { p.ID = id }
