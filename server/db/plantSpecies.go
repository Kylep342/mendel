package db

import (
	"database/sql"
	"time"

	"github.com/kylep342/mendel/model"
)

type PlantSpeciesTable struct {
	DB *sql.DB
}

// insert
func (repo *PlantSpeciesTable) Create(ps *model.PlantSpecies) error {
	query := `
		INSERT INTO plant_species (name, taxon)
		VALUES ($1, $2)
		returning id
	`
	err := repo.DB.QueryRow(query, ps.Name, ps.Taxon).Scan(&ps.Id, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// read all
func (repo *PlantSpeciesTable) GetAll() ([]model.PlantSpecies, error) {
	rows, err := repo.DB.Query(`SELECT id, name, taxon, created_at, updated_at FROM plant_species`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var speciesList []model.PlantSpecies
	for rows.Next() {
		var ps model.PlantSpecies
		if err := rows.Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		speciesList = append(speciesList, ps)
	}
	return speciesList, nil
}

// read one
func (repo *PlantSpeciesTable) GetByID(id string) (*model.PlantSpecies, error) {
	var ps model.PlantSpecies
	err := repo.DB.QueryRow(`
		SELECT id, name, taxon, created_at, updated_at FROM plant_species WHERE id = $1
	`, id).Scan(&ps.Id, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// update
func (repo *PlantSpeciesTable) Update(ps *model.PlantSpecies) error {
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
