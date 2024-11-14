-- name: CreateOrderItem :one
INSERT INTO orderitem (
                       order_id,
                       food_id,
                       quantity,
                       item_price)
VALUES (
        $1,
        $2,
        $3,
        $4)
    RETURNING *;

-- name: GetOrderItemByOrder :one
SELECT * FROM orderitem
WHERE order_id = $1
ORDER BY id;

-- name: UpdateOrderItem :one
UPDATE orderitem
SET quantity = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM orderitem
WHERE id = $1;