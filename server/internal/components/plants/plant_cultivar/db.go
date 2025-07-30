package plant_cultivar

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kylep342/mendel/internal/components/plants"
	"github.com/kylep342/mendel/internal/constants"
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

type Store struct {
	conn *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{conn: pool}
}

// Create inserts a new plant cultivar into the database
func (s *Store) Create(ctx context.Context, pc *plants.PlantCultivar) error {
	err := s.conn.QueryRow(ctx, queryCreatePlantCultivar, pc.SpeciesID, pc.Name, pc.Cultivar, pc.Genetics).Scan(&pc.ID, &pc.CreatedAt, &pc.UpdatedAt)
	return err
}

// GetAll retrieves all plant cultivars from the database
func (s *Store) GetAll(ctx context.Context) ([]plants.PlantCultivar, error) {
	rows, err := s.conn.Query(ctx, queryGetAllPlantCultivars)
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
func (s *Store) GetByID(ctx context.Context, id string) (plants.PlantCultivar, error) {
	var pc plants.PlantCultivar
	err := s.conn.QueryRow(ctx, queryGetPlantCultivarByID, id).Scan(&pc.ID, &pc.SpeciesID, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics)
	return pc, err
}

// Update modifies an existing plant cultivar in the database
func (s *Store) Update(ctx context.Context, pc *plants.PlantCultivar) error {
	err := s.conn.QueryRow(ctx, queryUpdatePlantCultivar, pc.ID, pc.SpeciesID, pc.Name, pc.Cultivar, pc.Genetics).Scan(
		&pc.ID, &pc.SpeciesID, &pc.Name, &pc.Cultivar, &pc.CreatedAt, &pc.UpdatedAt, &pc.Genetics,
	)
	return err
}

// Delete removes a plant cultivar from the database
func (s *Store) Delete(ctx context.Context, id string) error {
	_, err := s.conn.Exec(ctx, queryDeletePlantCultivar, id)
	return err
}
