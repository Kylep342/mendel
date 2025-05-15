package db

import (
	"database/sql"
	"time"

	"github.com/kylep342/mendel/constants"
	"github.com/kylep342/mendel/models"
)

const (
	// PlantTableName is the name of the plant table in the database
	TABLE_PLANT = constants.SchemaMendelCore + ".plant"

	// queryCreatePlant is the query template literal to create a new plant
	queryCreatePlant = `
		INSERT INTO ` + TABLE_PLANT + `
		(cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, planted_at
	`

	// queryGetAllPlants is the query template literal to get all plants
	queryGetAllPlants = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
		FROM ` + TABLE_PLANT

	// queryGetPlantByID is the query template literal to get a plant by ID
	queryGetPlantByID = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
		FROM ` + TABLE_PLANT + ` WHERE id = $1
	`

	// queryUpdatePlant is the query template literal to update a plant
	queryUpdatePlant = `
		UPDATE ` + TABLE_PLANT + `
		SET cultivar_id = $2, species_id = $3, seed_id = $4, pollen_id = $5, generation = $6, planted_at = $7, harvested_at = $8, genetics = $9, labels = $10
		WHERE id = $1
		RETURNING id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
	`

	// queryDeletePlant is the query template literal to delete a plant
	queryDeletePlant = `DELETE FROM ` + TABLE_PLANT + ` WHERE id = $1`
)

type PlantTable struct {
	DB *sql.DB
}

// GetAll retrieves all plants from the database
func (t *PlantTable) GetAll() ([]models.Plant, error) {
	rows, err := t.DB.Query(queryGetAllPlants)
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

// GetByID retrieves a plant by ID from the database
func (t *PlantTable) GetByID(id string) (models.Plant, error) {
	var ps models.Plant
	err := t.DB.QueryRow(
		queryGetPlantByID,
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

// Create saves a new plant to the database
func (t *PlantTable) Create(ps *models.Plant) error {
	err := t.DB.QueryRow(queryCreatePlant,
		ps.CultivarId,
		ps.SpeciesId,
		ps.SeedId,
		ps.PollenId,
		ps.Generation,
		time.Now(),
		nil,
		ps.Genetics,
		ps.Labels).Scan(
		&ps.Id,
		&ps.PlantedAt,
	)
	return err
}

// Update changes plants by IDs from the database
func (t *PlantTable) Update(ps *models.Plant) error {
	_, err := t.DB.Exec(queryUpdatePlant, ps.Id, ps.CultivarId, ps.SpeciesId, ps.SeedId, ps.PollenId, ps.Generation, ps.PlantedAt, ps.HarvestedAt, ps.Genetics, ps.Labels)
	return err
}

// Delete removes a plant by ID from the database
func (t *PlantTable) Delete(id string) error {
	_, err := t.DB.Exec(queryDeletePlant, id)
	return err
}
