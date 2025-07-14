package main

import (
	"github.com/ivcDark/newsbot/internal/config"
	"github.com/ivcDark/newsbot/internal/infrastructure/db"
	"github.com/ivcDark/newsbot/internal/parser"
	"log"
)

func main() {
	cfg := config.Load()

	log.Printf("Connecting to DB with driver=%s, dsn=%s", cfg.DBDriver, cfg.DBDSN)
	dbConn, err := db.Connect(cfg.DBDriver, cfg.DBDSN)
	if err != nil {
		log.Fatalf("Could not connect to %s due to %s", cfg.DBDriver, err)
	}
	log.Printf("Connected to DB: %s", cfg.DBDriver)

	err = db.RunMigrations("migrations", cfg.DBDSN)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Migrations applied")

	newsList, err := parser.ParseNews()
	if err != nil {
		log.Fatalf("Parse failed: %v", err)
	}

	repo := db.NewGormNewsRepository(dbConn)

	for _, item := range newsList {
		domainNews := item.ToDomain()
		existing, err := repo.FindByURL(domainNews.URL)
		if err != nil {
			log.Printf("DB error: %v", err)
			continue
		}
		if existing != nil {
			log.Printf("Already exists: %s", domainNews.URL)
			continue
		}
		err = repo.Save(domainNews)
		if err != nil {
			log.Printf("Error saving: %v", err)
		} else {
			log.Printf("Saved: %s", domainNews.Title)
		}
	}
}
