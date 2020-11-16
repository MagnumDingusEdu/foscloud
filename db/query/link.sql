-- name: CreateLink :one
INSERT INTO links (node, link, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLink :one
SELECT *
FROM links
WHERE id = $1
LIMIT 1;

-- name: ListLinks :many
SELECT *
FROM links
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateLink :one
UPDATE links
SET (link, password) = ($2, $3)
WHERE id = $1
RETURNING *;

-- name: DeleteLink :exec
DELETE
FROM links
WHERE id = $1;




