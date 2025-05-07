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
func (repo *PlantCultivarTable) Create(pc *models.PlantCultivar) error {
	query := `
		INSERT INTO plant_cultivar (species_id, name, cultivar, genetics)
		VALUES ($1, $2, $3, $4)
		returning id, created_at, updated_at
	`
	err := repo.DB.QueryRow(query, pc.SpeciesId, pc.Name, pc.Cultivar, pc.Genetics).Scan(&pc.Id, &pc.CreatedAt, &pc.UpdatedAt)
	return err
}

// read all
func (repo *PlantCultivarTable) GetAll() ([]models.PlantCultivar, error) {
	rows, err := repo.DB.Query(`SELECT id, species_id, name, cultivar, created_at, updated_at, genetics FROM plant_cultivar`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Cultivars []models.PlantCultivar
	for rows.Next() {
		var pc models.PlantCultivar
		if err := rows.Scan(&pc.Id, &pc.SpeciesId, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics); err != nil {
			return nil, err
		}
		Cultivars = append(Cultivars, pc)
	}
	return Cultivars, nil
}

// read one
func (repo *PlantCultivarTable) GetByID(id string) (*models.PlantCultivar, error) {
	var pc models.PlantCultivar
	err := repo.DB.QueryRow(`
		SELECT id, species_id, name, cultivar, created_at, updated_at, genetics FROM plant_cultivar WHERE id = $1
	`, id).Scan(&pc.Id, &pc.SpeciesId, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics)

	if err != nil {
		return nil, err
	}
	return &pc, nil
}

// update
func (repo *PlantCultivarTable) Update(pc *models.PlantCultivar) error {
	err := repo.DB.QueryRow(`
		UPDATE plant_cultivar
		SET species_id = $1, name = $2, cultivar = $3, genetics = $4
		WHERE id = $5
		RETURNING id, species_id, name, cultivar, created_at, updated_at, genetics
	`, pc.SpeciesId, pc.Name, pc.Cultivar, pc.Genetics, pc.Id).Scan(
		&pc.Id, &pc.SpeciesId, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics,
	)
	return err
}

// delete
func (repo *PlantCultivarTable) Delete(id string) error {
	_, err := repo.DB.Exec(`DELETE FROM plant_cultivar WHERE id = $1`, id)
	return err
}
