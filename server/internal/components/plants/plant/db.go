package plant

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kylep342/mendel/internal/constants"
)

// SQL queries for the plant table.
// Using constants for table names and queries keeps them organized and easy to modify.
const (
	tablePlant = constants.SchemaMendelCore + ".plant"

	queryCreatePlant = `
		INSERT INTO ` + tablePlant + ` (cultivar_id, species_id, seed_id, pollen_id, generation, created_at, updated_at, genetics, labels)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at`

	queryListPlants = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, created_at, updated_at, genetics, labels
		FROM ` + tablePlant

	queryGetPlantByID = `
		SELECT id, cultivar_id, species_id, seed_id, pollen_id, generation, created_at, updated_at, genetics, labels
		FROM ` + tablePlant + ` WHERE id = $1`

	queryUpdatePlant = `
		UPDATE ` + tablePlant + `
		SET cultivar_id = $2, species_id = $3, seed_id = $4, pollen_id = $5, generation = $6, created_at = $7, updated_at = $8, genetics = $9, labels = $10
		WHERE id = $1
		RETURNING id, cultivar_id, species_id, seed_id, pollen_id, generation, created_at, updated_at, genetics, labels`

	queryDeletePlant = `DELETE FROM ` + tablePlant + ` WHERE id = $1`
)

// Store handles all database operations for the Plant entity.
// It uses a pgxpool for connection management.
type Store struct {
	Conn *pgxpool.Pool
}

// NewStore creates a new Plant Store.
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{Conn: pool}
}

// List retrieves all plants from the database.
func (s *Store) GetAll(ctx context.Context) ([]Plant, error) {
	rows, err := s.Conn.Query(ctx, queryListPlants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var p Plant
		if err := rows.Scan(
			&p.ID,
			&p.CultivarID,
			&p.SpeciesID,
			&p.SeedID,
			&p.PollenID,
			&p.Generation,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Genetics,
			&p.Labels,
		); err != nil {
			return nil, err
		}
		plants = append(plants, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plants, nil
}

// GetByID retrieves a single plant by its ID.
func (s *Store) GetByID(ctx context.Context, id string) (Plant, error) {
	var p Plant
	err := s.Conn.QueryRow(ctx, queryGetPlantByID, id).Scan(
		&p.ID,
		&p.CultivarID,
		&p.SpeciesID,
		&p.SeedID,
		&p.PollenID,
		&p.Generation,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Genetics,
		&p.Labels,
	)
	return p, err
}

// Create inserts a new plant record into the database.
// It scans the RETURNING values back into the provided struct.
func (s *Store) Create(ctx context.Context, p *Plant) error {
	p.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err := s.Conn.QueryRow(ctx, queryCreatePlant,
		p.CultivarID,
		p.SpeciesID,
		p.SeedID,
		p.PollenID,
		p.Generation,
		p.CreatedAt,
		p.UpdatedAt,
		p.Genetics,
		p.Labels,
	).Scan(&p.ID, &p.CreatedAt)

	return err
}

// Update modifies an existing plant record.
// It scans the full updated record back into the provided struct.
func (s *Store) Update(ctx context.Context, p *Plant) error {
	return s.Conn.QueryRow(ctx, queryUpdatePlant,
		p.ID,
		p.CultivarID,
		p.SpeciesID,
		p.SeedID,
		p.PollenID,
		p.Generation,
		p.CreatedAt,
		p.UpdatedAt,
		p.Genetics,
		p.Labels,
	).Scan(
		&p.ID,
		&p.CultivarID,
		&p.SpeciesID,
		&p.SeedID,
		&p.PollenID,
		&p.Generation,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Genetics,
		&p.Labels,
	)
}

// Delete removes a plant record from the database by its ID.
func (s *Store) Delete(ctx context.Context, id string) error {
	_, err := s.Conn.Exec(ctx, queryDeletePlant, id)
	return err
}
