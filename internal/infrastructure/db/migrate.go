package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func RunMigrations(path string, dbFile string) error {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("Cannot open file sqlite db: %w", err)
	}
	defer db.Close()

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Cannot read migrations dir: %w", path, err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".sql" {
			continue
		}
		sqlContent, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return fmt.Errorf("Cannot read migration file %s: %w", file.Name(), err)
		}

		_, err = db.Exec(string(sqlContent))
		if err != nil {
			return fmt.Errorf("Failed to apply migration %s: %w", file.Name(), err)
		}
		fmt.Printf("Migration %s applied successfully", file.Name())
	}

	return nil
}
