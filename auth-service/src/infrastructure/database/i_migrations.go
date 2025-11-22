package database

type IMigrations interface {
	Up()
	Down()
}
