package plants

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

// idk man i'm just trying to make it do more than nothing
func (p *PlantSpecies) GetID() string { return p.ID }

// SetID sets the ID of the underlying record
//
//	probs should not exist; need to see where it's being called
func (p *PlantSpecies) SetID(id string) { p.ID = id }
