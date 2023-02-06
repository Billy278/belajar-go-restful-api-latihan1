package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicHandler(errorRollback)
		panic(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanicHandler(errorCommit)
	}
}
