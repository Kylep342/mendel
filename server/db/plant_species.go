package db

import (
	"database/sql"

	"github.com/kylep342/mendel/constants"
	"github.com/kylep342/mendel/models"
)

const (
	// PlantSpeciesTableName is the name of the plant species table in the database
	TABLE_PLANT_SPECIES = constants.SchemaMendelCore + ".plant_species"

	// queryCreatePlantSpecies is the query template literal to create a new plant species
	queryCreatePlantSpecies = `
		INSERT INTO ` + TABLE_PLANT_SPECIES + `
		(name, taxon)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	// queryGetAllPlantSpecies is the query template literal to get all plant species
	queryGetAllPlantSpecies = `
		SELECT id, name, taxon, created_at, updated_at
		FROM ` + TABLE_PLANT_SPECIES

	// queryGetByIDPlantSpecies is the query template literal to get a plant species by ID
	queryGetByIDPlantSpecies = `
		SELECT id, name, taxon, created_at, updated_at
		FROM ` + TABLE_PLANT_SPECIES + ` WHERE id = $1
	`

	// queryUpdatePlantSpecies is the query template literal to update a plant species
	queryUpdatePlantSpecies = `
		UPDATE ` + TABLE_PLANT_SPECIES + `
		SET name = $1, taxon = $2
		WHERE id = $3
		RETURNING id, name, taxon, created_at, updated_at
	`
	// queryDeletePlantSpecies is the query template literal to delete a plant species
	queryDeletePlantSpecies = `DELETE FROM ` + TABLE_PLANT_SPECIES + ` WHERE id = $1`
)

type PlantSpeciesTable struct {
	DB *sql.DB
}

// Create inserts a new plant species into the database
func (t *PlantSpeciesTable) Create(ps *models.PlantSpecies) error {
	query := queryCreatePlantSpecies
	err := t.DB.QueryRow(query, ps.Name, ps.Taxon).Scan(&ps.Id, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// GetAll retrieves all plant species from the database
func (t *PlantSpeciesTable) GetAll() ([]models.PlantSpecies, error) {
	rows, err := t.DB.Query(queryGetAllPlantSpecies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Species []models.PlantSpecies
	for rows.Next() {
		var ps models.PlantSpecies
		if err := rows.Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		Species = append(Species, ps)
	}
	return Species, nil
}

// GetByID retrieves a plant species identified by argument `id` from the database
func (t *PlantSpeciesTable) GetByID(id string) (models.PlantSpecies, error) {
	var ps models.PlantSpecies
	err := t.DB.QueryRow(queryGetByIDPlantSpecies, id).Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt)
	return ps, err
}

// Update updates a plant species identified by argument `id` in the database
func (t *PlantSpeciesTable) Update(ps *models.PlantSpecies) error {
	err := t.DB.QueryRow(queryUpdatePlantSpecies, ps.Name, ps.Taxon, ps.Id).Scan(
		&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt,
	)
	return err
}

// Delete removes a plant species identified by argument `id` from the database
func (t *PlantSpeciesTable) Delete(id string) error {
	_, err := t.DB.Exec(queryDeletePlantSpecies, id)
	return err
}
