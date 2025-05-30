package db

import (
	"context"
	"database/sql"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/models/plants"
)

const (
	// PlantCultivarTableName is the name of the table in the database
	tablePlantCultivar = constants.SchemaMendelCore + ".plant_cultivar"

	// queryCreatePlantCultivar is the query template literal to create a new plant cultivar
	queryCreatePlantCultivar = `
		INSERT INTO ` + tablePlantCultivar + `
		(species_id, name, cultivar, genetics)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	// queryGetAllPlantCultivars is the query template literal to get all plant cultivars
	queryGetAllPlantCultivars = `
		SELECT id, species_id, name, cultivar, created_at, updated_at, genetics
		FROM ` + tablePlantCultivar

	// queryGetPlantCultivarByID is the query template literal to get a plant cultivar by ID
	queryGetPlantCultivarByID = `
		SELECT id, species_id, name, cultivar, created_at, updated_at, genetics
		FROM ` + tablePlantCultivar + ` WHERE id = $1
	`
	// queryUpdatePlantCultivar is the query template literal to update a plant cultivar
	queryUpdatePlantCultivar = `
		UPDATE ` + tablePlantCultivar + `
		SET species_id = $2, name = $3, cultivar = $4, genetics = $5
		WHERE id = $1
		RETURNING id, species_id, name, cultivar, created_at, updated_at, genetics
	`
	// queryDeletePlantCultivar is the query template literal to delete a plant cultivar
	queryDeletePlantCultivar = `DELETE FROM ` + tablePlantCultivar + ` WHERE id = $1`
)

type PlantCultivarTable struct {
	DB *sql.DB
}

// Create inserts a new plant cultivar into the database
func (repo *PlantCultivarTable) Create(ctx context.Context, pc *plants.PlantCultivar) error {
	err := repo.DB.QueryRowContext(ctx, queryCreatePlantCultivar, pc.SpeciesID, pc.Name, pc.Cultivar, pc.Genetics).Scan(&pc.ID, &pc.CreatedAt, &pc.UpdatedAt)
	return err
}

// GetAll retrieves all plant cultivars from the database
func (repo *PlantCultivarTable) GetAll(ctx context.Context) ([]plants.PlantCultivar, error) {
	rows, err := repo.DB.QueryContext(ctx, queryGetAllPlantCultivars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Cultivars []plants.PlantCultivar
	for rows.Next() {
		var pc plants.PlantCultivar
		if err := rows.Scan(&pc.ID, &pc.SpeciesID, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics); err != nil {
			return nil, err
		}
		Cultivars = append(Cultivars, pc)
	}
	return Cultivars, nil
}

// GetByID retrieves a plant cultivar identified by arg `id` from the database
func (repo *PlantCultivarTable) GetByID(ctx context.Context, id string) (plants.PlantCultivar, error) {
	var pc plants.PlantCultivar
	err := repo.DB.QueryRowContext(ctx, queryGetPlantCultivarByID, id).Scan(&pc.ID, &pc.SpeciesID, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics)
	return pc, err
}

// Update modifies an existing plant cultivar in the database
func (repo *PlantCultivarTable) Update(ctx context.Context, pc *plants.PlantCultivar) error {
	err := repo.DB.QueryRowContext(ctx, queryUpdatePlantCultivar, pc.ID, pc.SpeciesID, pc.Name, pc.Cultivar, pc.Genetics).Scan(
		&pc.ID, &pc.SpeciesID, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics,
	)
	return err
}

// Delete removes a plant cultivar from the database
func (repo *PlantCultivarTable) Delete(ctx context.Context, id string) error {
	_, err := repo.DB.ExecContext(ctx, queryDeletePlantCultivar, id)
	return err
}
