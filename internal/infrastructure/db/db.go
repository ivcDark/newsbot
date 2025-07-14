package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(driver, dsn string) (*gorm.DB, error) {
	switch driver {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %s: %w", driver, err)
		}
		return db, nil
	default:
		return nil, fmt.Errorf("driver %s not supported", driver)
	}
}
