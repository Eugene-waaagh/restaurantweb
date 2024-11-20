-- name: AddCustomizationToOrderItem :one
INSERT INTO ordercustomization (order_item_id, food_customization_id)
VALUES ($1, $2)
    RETURNING *;

-- name: GetOrderItemCustomizations :many
SELECT fc.customization_type, fc.value
FROM ordercustomization oc
         JOIN foodcustomization fc ON oc.food_customization_id = fc.id
WHERE oc.order_item_id = $1;

-- name: DeleteOrderItemCustomization :exec
DELETE FROM ordercustomization
WHERE order_item_id = $1 AND food_customization_id = $2;
