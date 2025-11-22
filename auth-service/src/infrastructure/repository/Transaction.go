package repository

import "github.com/jmoiron/sqlx"

func (r *BaseRepository) WithTransaction(fn func(tx *sqlx.Tx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if rec := recover(); rec != nil {
			_ = tx.Rollback()
			panic(rec)
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
