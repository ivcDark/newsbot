package migrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations(db *sql.DB, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Failed to read directory %s: %s", dir, err)
	}

	var sqlFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}

	sort.Strings(sqlFiles)

	for _, file := range sqlFiles {
		path := filepath.Join(dir, file)
		fmt.Printf("Running migration: %s\n", path)

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("Failed to read file %s: %s", path, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("Failed to execute migration %s: %s", path, err)
		}
	}

	return nil
}
