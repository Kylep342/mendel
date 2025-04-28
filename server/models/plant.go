package models

import (
	"time"
)

type Plant struct {
	Id          string    `db:"id"`
	CultivarId  string    `db:"cultivar_id"`
	SpeciesId   string    `db:"species_id"`
	SeedId      string    `db:"seed_id"`
	PollenId    string    `db:"pollen_id"`
	Generation  uint32    `db:"generation"`
	PlantedAt   time.Time `db:"planted_at"`
	HarvestedAt time.Time `db:"harvested_at"`
	Genetics    struct{}  `db:"genetics"`
	Labels      struct{}  `db:"labels"`
}
