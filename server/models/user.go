package models

import (
	"time"
)

type User struct {
	ID          string    `db:"id"`
	Username    string    `db:"username"`
	Enabled     bool      `db:"enabled"`
	WebSettings struct{}  `db:"web_settings"`
	LastLogin   time.Time `db:"last_login"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
