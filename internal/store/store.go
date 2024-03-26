package store

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ConnectToDatabase() (*sql.DB, error) {
	database, err := sql.Open("sqlite", "../internal/store/store.db")

	if err != nil {
		return &sql.DB{}, err
	}

	return database, nil
}
