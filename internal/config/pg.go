package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	hostEnvName = "DB_HOST"
	userEnvName = "DB_USERNAME"
	passEnvName = "DB_PASSWORD"
	portEnvName = "DB_PORT"
	dbEnvName   = "DB_DATABASE"
)

// PGConfig defines the configuration interface for a PostgreSQL database.
// It provides a method to get the Data Source Name (DSN) as a string.
type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

// NewPGConfig creates and returns a new PGConfig instance.
func NewPGConfig() (PGConfig, error) {
	host := os.Getenv(hostEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg host not found")
	}

	user := os.Getenv(userEnvName)
	if len(user) == 0 {
		return nil, errors.New("pg username not found")
	}

	pass := os.Getenv(passEnvName)
	if len(pass) == 0 {
		return nil, errors.New("pg password not found")
	}

	port := os.Getenv(portEnvName)
	if len(port) == 0 {
		return nil, errors.New("pg port not found")
	}

	db := os.Getenv(dbEnvName)
	if len(db) == 0 {
		return nil, errors.New("pg database not found")
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, db, user, pass)

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
