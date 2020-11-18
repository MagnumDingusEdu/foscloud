-- name: CreateAuthToken :one
INSERT INTO authtoken (token, account)
VALUES ($1, $2)
RETURNING *;

-- name: GetAuthTokenByID :one
SELECT *
FROM authtoken
WHERE id = $1
LIMIT 1;

-- name: GetAuthTokenByAccount :one
SELECT *
FROM authtoken
WHERE account = $1
LIMIT 1;

-- name: UpdateAuthTokenValue :one
UPDATE authtoken
SET token = $2
WHERE account = $1
RETURNING *;


-- name: UpdateAuthTokenDate :one
UPDATE authtoken
SET last_used = $2
WHERE account = $1
RETURNING *;

-- name: DeleteAuthToken :exec
DELETE
FROM authtoken
WHERE account = $1;
