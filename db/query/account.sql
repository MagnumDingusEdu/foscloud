-- name: CreateAccount :one
INSERT INTO accounts (name, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET (name, username, email, password) = ($2, $3, $4, $5)
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1;

-- name: UpdateLastLogin :one
UPDATE accounts
SET last_login = $2
WHERE id = $1
RETURNING *;

-- name: CountAccounts :one
SELECT count(*)
FROM accounts;

-- name: CheckAccount :one
SELECT *
FROM accounts
WHERE (username = $1 OR email = $1)
LIMIT 1;