package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ivcDark/newsbot/internal/migrate"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Запустить миграции базы данных",
	Run: func(cmd *cobra.Command, args []string) {
		dbType := os.Getenv("DB_DRIVER")
		var db *sql.DB
		var err error
		var migrationsDir string

		switch dbType {
		case "sqlite":
			db, err = sql.Open("sqlite3", "./newsbot.db")
			migrationsDir = "migrations/sqlite"
		case "postgres":
			db, err = sql.Open("postgres", os.Getenv("DB_DSN"))
			migrationsDir = "migrations/postgres"
		default:
			log.Fatalf("Unsupported DB_TYPE: %s", dbType)
		}

		if err != nil {
			log.Fatalf("Failed to open db: %v", err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatalf("Cannot connect to DB: %v", err)
		}

		err = migrate.RunMigrations(db, migrationsDir)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		fmt.Println("Миграции успешно применены")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
