-- name: CreateNode :one
INSERT INTO nodes (parent_id, name, filesize, depth, lineage, owner)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetNode :one
SELECT *
FROM nodes
WHERE id = $1
LIMIT 1;

-- name: ListNodes :many
SELECT *
FROM nodes
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeleteNode :exec
DELETE
FROM nodes
WHERE id = $1;

-- name: CountNodes :one
SELECT count(*) FROM nodes;




