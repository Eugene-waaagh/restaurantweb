-- name: CreateCategory :one
INSERT INTO foodcatalogue (
    name
) VALUES (
             $1
         ) RETURNING *;

-- name: GetCategory :one
SELECT * FROM foodcatalogue
WHERE id = $1 LIMIT 1;

-- name: ListCategory :many
SELECT * FROM food
ORDER BY id;

-- name: UpdateCategory :one
UPDATE foodcatalogue
SET name = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM foodcatalogue
WHERE id = $1;
