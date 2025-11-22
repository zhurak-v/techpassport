package database

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	dbUrl         string
	migrationPath string
}

func NewMigration() IMigrations {
	return &Migration{
		dbUrl:         "postgres://postgres:111@postgres/techpassport_auth_db?sslmode=disable",
		migrationPath: "file:///app/migrations",
	}
}

func (m *Migration) Up() {
	migrator, err := migrate.New(m.migrationPath, m.dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}

func (m *Migration) Down() {
	migrator, err := migrate.New(m.migrationPath, m.dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Migration rolled back successfully")
}
