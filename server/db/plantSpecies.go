package db

import (
	"database/sql"
	"time"

	"github.com/kylep342/mendel/models"
)

type PlantSpeciesTable struct {
	DB *sql.DB
}

func NewPlantSpeciesTable(db *sql.DB) *PlantSpeciesTable {
	return &PlantSpeciesTable{
		DB: db,
	}
}

// insert
func (repo *PlantSpeciesTable) Create(ps *models.PlantSpecies) error {
	query := `
		INSERT INTO plant_species (name, taxon)
		VALUES ($1, $2)
		returning id, created_at, updated_at
	`
	err := repo.DB.QueryRow(query, ps.Name, ps.Taxon).Scan(&ps.Id, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// read all
func (repo *PlantSpeciesTable) GetAll() ([]models.PlantSpecies, error) {
	rows, err := repo.DB.Query(`SELECT id, name, taxon, created_at, updated_at FROM plant_species`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var speciesList []models.PlantSpecies
	for rows.Next() {
		var ps models.PlantSpecies
		if err := rows.Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		speciesList = append(speciesList, ps)
	}
	return speciesList, nil
}

// read one
func (repo *PlantSpeciesTable) GetByID(id string) (*models.PlantSpecies, error) {
	var ps models.PlantSpecies
	err := repo.DB.QueryRow(`
		SELECT id, name, taxon, created_at, updated_at FROM plant_species WHERE id = $1
	`, id).Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// update
func (repo *PlantSpeciesTable) Update(ps *models.PlantSpecies) error {
	_, err := repo.DB.Exec(`
		UPDATE plant_species
		SET name = $1, taxon = $2, updated_at = $3
		WHERE id = $4
	`, ps.Name, ps.Taxon, time.Now(), ps.Id)
	return err
}

// delete
func (repo *PlantSpeciesTable) Delete(id string) error {
	_, err := repo.DB.Exec(`DELETE FROM plant_species WHERE id = $1`, id)
	return err
}
