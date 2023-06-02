package helper

import (
	"database/sql"
)

func CommitOrRollback(err error, tx *sql.Tx) error {
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
