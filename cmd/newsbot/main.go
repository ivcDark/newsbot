package main

import (
	"github.com/ivcDark/newsbot/internal/config"
	"github.com/ivcDark/newsbot/internal/infrastructure/db"
	"log"
)

func main() {
	cfg := config.Load()
	log.Printf("Connecting to DB with driver=%s, dsn=%s", cfg.DBDriver, cfg.DBDSN)
	_, err := db.Connect(cfg.DBDriver, cfg.DBDSN)
	if err != nil {
		log.Fatalf("Could not connect to %s due to %s", cfg.DBDriver, err)
	}

	log.Printf("Connected to DB: %s", cfg.DBDriver)
}
