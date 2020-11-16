// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
	"time"
)

const countAccounts = `-- name: CountAccounts :one
SELECT count(*) FROM accounts
`

func (q *Queries) CountAccounts(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAccounts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (name, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id, name, username, email, password, created_at, last_login
`

type CreateAccountParams struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, username, email, password, created_at, last_login
FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, name, username, email, password, created_at, last_login
FROM accounts
WHERE id = $1
LIMIT 1 FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, name, username, email, password, created_at, last_login
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.LastLogin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET (name, username, email, password) = ($2, $3, $4, $5)
WHERE id = $1
RETURNING id, name, username, email, password, created_at, last_login
`

type UpdateAccountParams struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const updateLastLogin = `-- name: UpdateLastLogin :one
UPDATE accounts
SET last_login = $2
WHERE id = $1
RETURNING id, name, username, email, password, created_at, last_login
`

type UpdateLastLoginParams struct {
	ID        int64     `json:"id"`
	LastLogin time.Time `json:"last_login"`
}

func (q *Queries) UpdateLastLogin(ctx context.Context, arg UpdateLastLoginParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateLastLogin, arg.ID, arg.LastLogin)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}
