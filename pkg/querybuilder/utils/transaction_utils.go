package utils

import (
	"database/sql"

	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

func Begin(databaseContext *structs.DataBaseContext) (*structs.Transaction, error) {
	var err error = nil
	var tx *structs.Transaction

	switch databaseContext.Type {
	case types.SqlLite, types.Mysql, types.Postgres:
		db := databaseContext.DB.(*sql.DB)
		txDb, err := db.Begin()

		if err == nil {
			tx = &structs.Transaction{Type: databaseContext.Type, TransactionContext: txDb}
		}

		break
	}

	return tx, err
}

func Commit(transaction *structs.Transaction) error {
	var err error = nil

	switch transaction.Type {
	case types.SqlLite, types.Mysql, types.Postgres:
		tx := transaction.TransactionContext.(*sql.Tx)
		err = tx.Commit()
		break
	}

	return err
}

func Rollback(transaction *structs.Transaction) error {
	var err error = nil

	switch transaction.Type {
	case types.SqlLite, types.Mysql, types.Postgres:
		tx := transaction.TransactionContext.(*sql.Tx)
		err = tx.Rollback()
		break
	}

	return err
}
