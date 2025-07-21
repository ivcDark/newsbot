package main

import (
	"database/sql"
	"fmt"
	"github.com/ivcDark/newsbot/internal/migrate"
	"github.com/ivcDark/newsbot/internal/repository"
	"github.com/ivcDark/newsbot/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./newsbot.db")
	if err != nil {
		log.Fatal("Failed to open db: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close db: %v", err)
		}
	}(db)

	err = migrate.RunMigrations(db, "migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Successfully run migrations")

	newsRepo := repository.NewSQLiteNewsRepository(db)
	var _ repository.NewsRepository = newsRepo
	newsService := service.NewNewsService(newsRepo)

	err = newsService.FetchAndSaveNews("https://63.ru/text/")

	if err != nil {
		log.Fatalf("Failed to fetch and save news: %v", err)
	}
}
