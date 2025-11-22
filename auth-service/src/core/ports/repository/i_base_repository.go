package repository

import "github.com/jmoiron/sqlx"

type BaseRepositoryReader interface{}
type BaseRepositoryWriter interface {
	WithTransaction(fn func(tx *sqlx.Tx) error) error
}
type BaseRepositoryContract interface {
	BaseRepositoryReader
	BaseRepositoryWriter
}
