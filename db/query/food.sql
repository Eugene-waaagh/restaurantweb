-- name: CreateFood :one
INSERT INTO food (
    name,
    description,
    price,
    category_id
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

DELETE FROM food
WHERE name = $1;