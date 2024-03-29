package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	fail := func(err error) error {
		return fmt.Errorf("CreateOrder: %v", err)
	}

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer tx.Rollback()

	q := New(tx)
	err = fn(q)
	if err != nil {
		fail(err)
	}

	return tx.Commit()
}
