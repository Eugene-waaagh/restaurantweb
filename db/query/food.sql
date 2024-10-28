-- name: CreateCategory :one
INSERT INTO food_catalogue (
    name
) VALUES (
             $1
         ) RETURNING *;

-- name: CreateFood :one
INSERT INTO food (
    name,
    description,
    price,
    category_id
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: UpdateCategory :one
UPDATE food_catalogue
SET name = $2
WHERE id = $1
    RETURNING *;



-- name: DeleteCategory :exec
DELETE FROM food_catalogue
WHERE name = $1;

DELETE FROM food
WHERE name = $1;