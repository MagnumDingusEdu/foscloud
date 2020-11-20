package db

import (
	"context"
	"database/sql"
	"fmt"
	"foscloud/utils"
	"time"
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

type IncorrectCredentialsError struct {
	err error
}

func (i IncorrectCredentialsError) Error() string {
	return fmt.Sprintf("The given credentials are incorrect.")
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

type LoginAccountTxParams struct {
	LoginID  string `json:"username"`
	Password string `json:"password"`
}

type LoginAccountTxResult struct {
	Account   Account   `json:"account"`
	Authtoken Authtoken `json:"authtoken"`
}

func (store *Store) LoginAccountTx(ctx context.Context, arg LoginAccountTxParams) (LoginAccountTxResult, error) {
	var result LoginAccountTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Account, err = store.CheckAccount(ctx, arg.LoginID)
		if err != nil {
			if err == sql.ErrNoRows {
				if err, ok := err.(*IncorrectCredentialsError); ok {
					return err
				} else {
					return err
				}
			}
			return err
		}
		_, err = utils.CheckPassword(result.Account.Password, []byte(arg.Password))
		if err != nil {
			if err, ok := err.(*IncorrectCredentialsError); ok {
				return err
			} else {
				return err
			}

		}
		authtoken, err := store.GetAuthTokenByAccount(ctx, result.Account.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.Authtoken, err = store.CreateAuthToken(ctx, CreateAuthTokenParams{
					Token:   utils.GenerateAuthToken(),
					Account: result.Account.ID,
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}

		} else {
			result.Authtoken = authtoken
		}
		_, err = store.UpdateLastLogin(ctx, UpdateLastLoginParams{
			ID:        result.Account.ID,
			LastLogin: time.Now(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
