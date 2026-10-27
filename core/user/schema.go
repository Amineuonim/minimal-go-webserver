package user

import "github.com/jmoiron/sqlx"

func EnsureSchema(db *sqlx.DB) error {
	const schema = `
CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);`

	_, err := db.Exec(schema)
	return err
}
