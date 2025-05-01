package db

import (
	"database/sql"

	"github.com/kylep342/mendel/models"
)

type PlantCultivarTable struct {
	DB *sql.DB
}

func NewPlantCultivarTable(db *sql.DB) *PlantCultivarTable {
	return &PlantCultivarTable{
		DB: db,
	}
}

// insert
func (repo *PlantCultivarTable) Create(ps *models.PlantCultivar) error {
	query := `
		INSERT INTO plant_cultivar (species_id, cultivar, genetics)
		VALUES ($1, $2, $3)
		returning id, created_at, updated_at
	`
	err := repo.DB.QueryRow(query, ps.SpeciesId, ps.Cultivar, ps.Genetics).Scan(&ps.Id, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// read all
func (repo *PlantCultivarTable) GetAll() ([]models.PlantCultivar, error) {
	rows, err := repo.DB.Query(`SELECT id, species_id, cultivar, created_at, updated_at, genetics FROM plant_cultivar`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var CultivarList []models.PlantCultivar
	for rows.Next() {
		var ps models.PlantCultivar
		if err := rows.Scan(&ps.Id, &ps.SpeciesId, &ps.Cultivar, &ps.CreatedAt, &ps.UpdatedAt, &ps.Genetics); err != nil {
			return nil, err
		}
		CultivarList = append(CultivarList, ps)
	}
	return CultivarList, nil
}

// read one
func (repo *PlantCultivarTable) GetByID(id string) (*models.PlantCultivar, error) {
	var ps models.PlantCultivar
	err := repo.DB.QueryRow(`
		SELECT id, species_id, cultivar, created_at, updated_at, genetics FROM plant_cultivar WHERE id = $1
	`, id).Scan(&ps.Id, &ps.SpeciesId, &ps.Cultivar, &ps.CreatedAt, &ps.UpdatedAt, &ps.Genetics)

	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// update
func (repo *PlantCultivarTable) Update(ps *models.PlantCultivar) error {
	err := repo.DB.QueryRow(`
		UPDATE plant_cultivar
		SET species_id = $1, cultivar = $2, genetics = $3
		WHERE id = $4
		RETURNING id, species_id, cultivar, created_at, updated_at, genetics
	`, ps.SpeciesId, ps.Cultivar, ps.Genetics, ps.Id).Scan(
		&ps.Id, &ps.SpeciesId, &ps.Cultivar, &ps.CreatedAt, &ps.UpdatedAt, &ps.Genetics,
	)
	return err
}

// delete
func (repo *PlantCultivarTable) Delete(id string) error {
	_, err := repo.DB.Exec(`DELETE FROM plant_cultivar WHERE id = $1`, id)
	return err
}
