-- name: CreateFood :one
INSERT INTO food (
    name,
    description,
    price,
    category_id
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: ListFood :many
SELECT * FROM food
WHERE category_id = $1
ORDER BY id;


DELETE FROM food
WHERE name = $1;