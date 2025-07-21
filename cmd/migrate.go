package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ivcDark/newsbot/internal/migrate"
	"github.com/spf13/cobra"
	"log"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Запустить миграции базы данных",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./newsbot.db")
		if err != nil {
			log.Fatalf("Failed to open db: %v", err)
		}
		defer db.Close()

		err = migrate.RunMigrations(db, "migrations")
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		fmt.Println("Миграции успешно применены")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
