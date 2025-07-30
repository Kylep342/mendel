package plant_species

import (
	"time"
)

type PlantSpecies struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Taxon     string    `db:"taxon" json:"taxon"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (p *PlantSpecies) GetID() string { return p.ID }

func (p *PlantSpecies) SetID(id string) { p.ID = id }
