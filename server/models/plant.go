package models

import (
	"time"
)

type Plant struct {
	Id          string      `db:"id" json:"id"`
	CultivarId  string      `db:"cultivar_id" json:"cultivar_id"`
	SpeciesId   string      `db:"species_id" json:"species_id"`
	SeedId      string      `db:"seed_id" json:"seed_id"`
	PollenId    string      `db:"pollen_id" json:"pollen_id"`
	Generation  uint32      `db:"generation" json:"generation"`
	PlantedAt   time.Time   `db:"planted_at" json:"planted_at"`
	HarvestedAt time.Time   `db:"harvested_at" json:"harvested_at"`
	Genetics    interface{} `db:"genetics" json:"genetics"`
	Labels      interface{} `db:"labels" json:"labels"`
}

func (p *Plant) GetID() string { return p.Id }

func (p *Plant) SetID(id string) { p.Id = id }
