-- name: CreateFoodCustomization :one
INSERT INTO foodcustomization (
    food_id,
    customization_type,
    value)
VALUES (
           $1,
           $2,
           $3)
    RETURNING *;

-- name: ListCustomizationsByFood :many
SELECT * FROM foodcustomization
WHERE food_id = $1
ORDER BY id;

-- name: UpdateFoodCustomization :one
UPDATE foodcustomization
SET customization_type = $2, value = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteFoodCustomization :exec
DELETE FROM foodcustomization
WHERE id = $1;