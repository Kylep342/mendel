package db

import (
	"database/sql"

	"github.com/kylep342/mendel/models"
)

const (
	// PlantCultivarTableName is the name of the table in the database
	TABLE_PLANT_CULTIVAR = "plant_cultivar"

	// queryCreatePlantCultivar is the query template literal to create a new plant cultivar
	queryCreatePlantCultivar = `
		INSERT INTO ` + TABLE_PLANT_CULTIVAR + `(species_id, name, cultivar, genetics)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	// queryGetAllPlantCultivars is the query template literal to get all plant cultivars
	queryGetAllPlantCultivars = `
		SELECT id, species_id, name, cultivar, created_at, updated_at, genetics
		FROM ` + TABLE_PLANT_CULTIVAR

	// queryGetPlantCultivarByID is the query template literal to get a plant cultivar by ID
	queryGetPlantCultivarByID = `
		SELECT id, species_id, name, cultivar, created_at, updated_at, genetics
		FROM ` + TABLE_PLANT_CULTIVAR + ` WHERE id = $1`
	// queryUpdatePlantCultivar is the query template literal to update a plant cultivar
	queryUpdatePlantCultivar = `
		UPDATE ` + TABLE_PLANT_CULTIVAR + `
		SET species_id = $2, name = $3, cultivar = $4, genetics = $5
		WHERE id = $1
		RETURNING id, species_id, name, cultivar, created_at, updated_at, genetics
	`
	// queryDeletePlantCultivar is the query template literal to delete a plant cultivar
	queryDeletePlantCultivar = `DELETE FROM ` + TABLE_PLANT_CULTIVAR + ` WHERE id = $1`
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
	err := repo.DB.QueryRow(queryCreatePlantCultivar, pc.SpeciesId, pc.Name, pc.Cultivar, pc.Genetics).Scan(&pc.Id, &pc.CreatedAt, &pc.UpdatedAt)
	return err
}

// read all
func (repo *PlantCultivarTable) GetAll() ([]models.PlantCultivar, error) {
	rows, err := repo.DB.Query(queryGetAllPlantCultivars)
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
	err := repo.DB.QueryRow(queryGetPlantCultivarByID, id).Scan(&pc.Id, &pc.SpeciesId, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics)

	if err != nil {
		return nil, err
	}
	return &pc, nil
}

// update
func (repo *PlantCultivarTable) Update(pc *models.PlantCultivar) error {
	err := repo.DB.QueryRow(queryUpdatePlantCultivar, pc.Id, pc.SpeciesId, pc.Name, pc.Cultivar, pc.Genetics).Scan(
		&pc.Id, &pc.SpeciesId, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics,
	)
	return err
}

// delete
func (repo *PlantCultivarTable) Delete(id string) error {
	_, err := repo.DB.Exec(queryDeletePlantCultivar, id)
	return err
}
