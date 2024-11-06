-- name: CreateFood :one
INSERT INTO food (
    name,
    description,
    price,
    category_id
) VALUES (
             $1,
             $2,
             $3,
             $4
         ) RETURNING *;

-- name: GetFood :one
SELECT * FROM food
WHERE id = $1 LIMIT 1;

-- name: ListFood :many
SELECT * FROM food
WHERE category_id = $1
ORDER BY id;

-- name: UpdateFood :one
UPDATE food
SET name = $2, description = $3, price = $4, category_id = $5
WHERE id = $1
RETURNING *;

DELETE FROM food
WHERE id = $1;