package repository

import (
	"database/sql"
	"fmt"
)

func NewRepository(driver string, db *sql.DB) (NewsRepository, error) {
	switch driver {
	case "sqlite3":
		return NewSQLiteNewsRepository(db), nil
	case "postgres":
		return NewPostgresNewsRepository(db), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}
}
