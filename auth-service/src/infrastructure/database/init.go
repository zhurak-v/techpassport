package database

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func InitDatabase() *sqlx.DB {
	dbURL := "postgres://postgres:111@postgres/techpassport_auth_db?sslmode=disable"

	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 10)

	return db
}
