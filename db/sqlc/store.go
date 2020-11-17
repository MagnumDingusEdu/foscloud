package db

import (
	"context"
	"database/sql"
	"fmt"
	"foscloud/utils"
)

// store provides all functions
//to run database queries and their combinations
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// Executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rBackErr := tx.Rollback(); rBackErr != nil {
			return fmt.Errorf("tx err : %v ; rollback err: %v", err, rBackErr)
		} else {
			return err
		}

	}
	return tx.Commit()
}

type RegisterTxParams struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (store *Store) RegisterAccount(ctx context.Context, arg RegisterTxParams) (Account, error) {
	var account Account
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		account, err = queries.CreateAccount(ctx, CreateAccountParams{
			Name:     arg.Name,
			Username: arg.Username,
			Email:    arg.Email,
			Password: utils.HashAndSalt(arg.Password),
		})
		if err != nil {
			return err
		}
		return nil
	})
	return account, err
}
