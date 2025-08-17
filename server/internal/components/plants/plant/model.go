package plant

import (
	"database/sql"
)

type Plant struct {
	ID         string       `db:"id" json:"id"`
	CultivarID string       `db:"cultivar_id" json:"cultivar_id"`
	SpeciesID  string       `db:"species_id" json:"species_id"`
	SeedID     string       `db:"seed_id" json:"seed_id"`
	PollenID   string       `db:"pollen_id" json:"pollen_id"`
	Generation uint32       `db:"generation" json:"generation"`
	CreatedAt  sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updated_at"`
	Genetics   interface{}  `db:"genetics" json:"genetics"`
	Labels     interface{}  `db:"labels" json:"labels"`
}

func (p *Plant) GetID() string { return p.ID }

func (p *Plant) SetID(id string) { p.ID = id }
