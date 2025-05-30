package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/models/plants"
)

const (
	// PlantTableName is the name of the plant table in the database
	tablePlant = constants.SchemaMendelCore + ".plant"

	// queryCreatePlant is the query template literal to create a new plant
	queryCreatePlant = `
		INSERT INTO ` + tablePlant + `
		(cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, planted_at
	`

	// queryGetAllPlants is the query template literal to get all plants
	queryGetAllPlants = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
		FROM ` + tablePlant

	// queryGetPlantByID is the query template literal to get a plant by ID
	queryGetPlantByID = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
		FROM ` + tablePlant + ` WHERE id = $1
	`

	// queryUpdatePlant is the query template literal to update a plant
	queryUpdatePlant = `
		UPDATE ` + tablePlant + `
		SET cultivar_id = $2, species_id = $3, seed_id = $4, pollen_id = $5, generation = $6, planted_at = $7, harvested_at = $8, genetics = $9, labels = $10
		WHERE id = $1
		RETURNING id, cultivar_id, species_id, seed_id, pollen_id, generation, planted_at, harvested_at, genetics, labels
	`

	// queryDeletePlant is the query template literal to delete a plant
	queryDeletePlant = `DELETE FROM ` + tablePlant + ` WHERE id = $1`
)

type PlantTable struct {
	DB *sql.DB
}

// GetAll retrieves all plants from the database
func (t *PlantTable) GetAll(ctx context.Context) ([]plants.Plant, error) {
	rows, err := t.DB.QueryContext(ctx, queryGetAllPlants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []plants.Plant
	for rows.Next() {
		var ps plants.Plant
		if err := rows.Scan(
			&ps.ID,
			&ps.CultivarID,
			&ps.SpeciesID,
			&ps.SeedID,
			&ps.PollenID,
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
func (t *PlantTable) GetByID(ctx context.Context, id string) (plants.Plant, error) {
	var ps plants.Plant
	err := t.DB.QueryRowContext(
		ctx,
		queryGetPlantByID,
		id).Scan(
		&ps.ID,
		&ps.CultivarID,
		&ps.SpeciesID,
		&ps.SeedID,
		&ps.PollenID,
		&ps.Generation,
		&ps.PlantedAt,
		&ps.HarvestedAt,
		&ps.Genetics,
		&ps.Labels)
	return ps, err
}

// Create saves a new plant to the database
func (t *PlantTable) Create(ctx context.Context, ps *plants.Plant) error {
	err := t.DB.QueryRowContext(
		ctx,
		queryCreatePlant,
		ps.CultivarID,
		ps.SpeciesID,
		ps.SeedID,
		ps.PollenID,
		ps.Generation,
		time.Now(),
		nil,
		ps.Genetics,
		ps.Labels).Scan(
		&ps.ID,
		&ps.PlantedAt,
	)
	return err
}

// Update changes plants by IDs from the database
func (t *PlantTable) Update(ctx context.Context, ps *plants.Plant) error {
	_, err := t.DB.ExecContext(ctx, queryUpdatePlant, ps.ID, ps.CultivarID, ps.SpeciesID, ps.SeedID, ps.PollenID, ps.Generation, ps.PlantedAt, ps.HarvestedAt, ps.Genetics, ps.Labels)
	return err
}

// Delete removes a plant by ID from the database
func (t *PlantTable) Delete(ctx context.Context, id string) error {
	_, err := t.DB.ExecContext(ctx, queryDeletePlant, id)
	return err
}
