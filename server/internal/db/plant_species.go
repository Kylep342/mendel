package db

import (
	"context"
	"database/sql"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/models/plants"
)

const (
	// PlantSpeciesTableName is the name of the plant species table in the database
	tablePlantSpecies = constants.SchemaMendelCore + ".plant_species"

	// queryCreatePlantSpecies is the query template literal to create a new plant species
	queryCreatePlantSpecies = `
		INSERT INTO ` + tablePlantSpecies + `
		(name, taxon)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	// queryGetAllPlantSpecies is the query template literal to get all plant species
	queryGetAllPlantSpecies = `
		SELECT id, name, taxon, created_at, updated_at
		FROM ` + tablePlantSpecies

	// queryGetByIDPlantSpecies is the query template literal to get a plant species by ID
	queryGetByIDPlantSpecies = `
		SELECT id, name, taxon, created_at, updated_at
		FROM ` + tablePlantSpecies + ` WHERE id = $1
	`

	// queryUpdatePlantSpecies is the query template literal to update a plant species
	queryUpdatePlantSpecies = `
		UPDATE ` + tablePlantSpecies + `
		SET name = $1, taxon = $2
		WHERE id = $3
		RETURNING id, name, taxon, created_at, updated_at
	`
	// queryDeletePlantSpecies is the query template literal to delete a plant species
	queryDeletePlantSpecies = `DELETE FROM ` + tablePlantSpecies + ` WHERE id = $1`
)

type PlantSpeciesTable struct {
	DB *sql.DB
}

// Create inserts a new plant species into the database
func (t *PlantSpeciesTable) Create(ctx context.Context, ps *plants.PlantSpecies) error {
	query := queryCreatePlantSpecies
	err := t.DB.QueryRowContext(ctx, query, ps.Name, ps.Taxon).Scan(&ps.ID, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// GetAll retrieves all plant species from the database
func (t *PlantSpeciesTable) GetAll(ctx context.Context) ([]plants.PlantSpecies, error) {
	rows, err := t.DB.QueryContext(ctx, queryGetAllPlantSpecies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Species []plants.PlantSpecies
	for rows.Next() {
		var ps plants.PlantSpecies
		if err := rows.Scan(&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		Species = append(Species, ps)
	}
	return Species, nil
}

// GetByID retrieves a plant species identified by argument `id` from the database
func (t *PlantSpeciesTable) GetByID(ctx context.Context, id string) (plants.PlantSpecies, error) {
	var ps plants.PlantSpecies
	err := t.DB.QueryRowContext(ctx, queryGetByIDPlantSpecies, id).Scan(&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt)
	return ps, err
}

// Update updates a plant species identified by argument `id` in the database
func (t *PlantSpeciesTable) Update(ctx context.Context, ps *plants.PlantSpecies) error {
	err := t.DB.QueryRowContext(ctx, queryUpdatePlantSpecies, ps.Name, ps.Taxon, ps.ID).Scan(
		&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt,
	)
	return err
}

// Delete removes a plant species identified by argument `id` from the database
func (t *PlantSpeciesTable) Delete(ctx context.Context, id string) error {
	_, err := t.DB.ExecContext(ctx, queryDeletePlantSpecies, id)
	return err
}
