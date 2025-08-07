package plant_species

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kylep342/mendel/internal/constants"
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

type Store struct {
	Conn *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{Conn: pool}
}

// Create inserts a new plant species into the database
func (s *Store) Create(ctx context.Context, ps *PlantSpecies) error {
	query := queryCreatePlantSpecies
	err := s.Conn.QueryRow(ctx, query, ps.Name, ps.Taxon).Scan(&ps.ID, &ps.CreatedAt, &ps.UpdatedAt)
	return err
}

// GetAll retrieves all plant species from the database
func (s *Store) GetAll(ctx context.Context) ([]PlantSpecies, error) {
	rows, err := s.Conn.Query(ctx, queryGetAllPlantSpecies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Species []PlantSpecies
	for rows.Next() {
		var ps PlantSpecies
		if err := rows.Scan(&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		Species = append(Species, ps)
	}
	return Species, nil
}

// GetByID retrieves a plant species identified by argument `id` from the database
func (s *Store) GetByID(ctx context.Context, id string) (PlantSpecies, error) {
	var ps PlantSpecies
	err := s.Conn.QueryRow(ctx, queryGetByIDPlantSpecies, id).Scan(&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt)
	return ps, err
}

// Update updates a plant species identified by argument `id` in the database
func (s *Store) Update(ctx context.Context, ps *PlantSpecies) error {
	err := s.Conn.QueryRow(ctx, queryUpdatePlantSpecies, ps.Name, ps.Taxon, ps.ID).Scan(
		&ps.ID, &ps.Name, &ps.Taxon, &ps.CreatedAt, &ps.UpdatedAt,
	)
	return err
}

// Delete removes a plant species identified by argument `id` from the database
func (s *Store) Delete(ctx context.Context, id string) error {
	_, err := s.Conn.Exec(ctx, queryDeletePlantSpecies, id)
	return err
}
