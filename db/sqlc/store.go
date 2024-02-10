//!TRANSACTIONS 

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// generating transactions
func (store *Store) execTX(ctx context.Context, function func(*Queries) error) error {
	TX, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("transaction has error and not started...\nerror is = %v", err)
	}

	dbquery := New(TX) // start query
	err = function(dbquery)
	if err != nil {
		if rlerr := TX.Rollback(); rlerr == err {
			return fmt.Errorf("transaction error %v rollback error %v", err, rlerr)
		}
		return fmt.Errorf("queries are not started so the transaction is rollback\nerror is = %v", err)
	}
	return TX.Commit()
}
