package db

import (
	"database/sql"
	"time"

	"github.com/kylep342/mendel/models"
)

type PlantTable struct {
	DB *sql.DB
}

// func NewPlantTable(db *sql.DB) *PlantTable {
// 	return &PlantTable{
// 		DB: db,
// 	}
// }

func (t *PlantTable) GetAll() ([]models.Plant, error) {
	rows, err := t.DB.Query("SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels FROM plant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Plant
	for rows.Next() {
		var ps models.Plant
		if err := rows.Scan(
			&ps.Id,
			&ps.CultivarId,
			&ps.SpeciesId,
			&ps.SeedId,
			&ps.PollenId,
			&ps.Generation,
			&ps.PlantedAt,
			&ps.HarvestedAt,
			&ps.Genetics,
			&ps.Labels); err != nil {
			return nil, err
		}
		result = append(result, ps)
	}
	return result, nil
}

func (t *PlantTable) GetByID(id string) (models.Plant, error) {
	var ps models.Plant
	err := t.DB.QueryRow(
		"SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels FROM plant WHERE id = $1",
		id).Scan(
		&ps.Id,
		&ps.CultivarId,
		&ps.SpeciesId,
		&ps.SeedId,
		&ps.PollenId,
		&ps.Generation,
		&ps.PlantedAt,
		&ps.HarvestedAt,
		&ps.Genetics,
		&ps.Labels)
	return ps, err
}

func (t *PlantTable) Create(ps *models.Plant) error {
	_, err := t.DB.Exec("INSERT INTO plant (cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels) VALUES ($1, $2, $3, $4, $5, $5, $7, $8, $9)", ps.CultivarId, ps.SpeciesId, ps.SeedId, ps.PollenId, ps.Generation, time.Now(), nil, ps.Genetics, ps.Labels)
	return err
}

func (t *PlantTable) Update(ps *models.Plant) error {
	_, err := t.DB.Exec("UPDATE plant SET cultivar_id = $2, species_id = $3, seed_id = $4, pollen_id = $5, generation = $6, planted_at = $7, harvested_at = $8, genetics = $9, labels = $10 WHERE id = $1", ps.Id, ps.CultivarId, ps.SpeciesId, ps.SeedId, ps.PollenId, ps.Generation, ps.PlantedAt, ps.HarvestedAt, ps.Genetics, ps.Labels)
	return err
}

func (t *PlantTable) Delete(id string) error {
	_, err := t.DB.Exec("DELETE FROM plant WHERE id = $1", id)
	return err
}
