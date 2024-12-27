-- name: CreateOrder :one
INSERT INTO "order" (
                     order_date,
                     total_price,
                     status)
VALUES (
        $1,
        $2,
        $3)
    RETURNING *;

-- name: GetOrderByStatus :one
SELECT * FROM "order"
WHERE status = $1
ORDER BY id;

-- name: UpdateOrderStatus :one
UPDATE "order"
SET status = $2
WHERE id = $1
    RETURNING *;

-- Special action

-- name: GetFullOrderDetails :many
SELECT
    o.id AS order_id,
    o.order_date,
    o.total_price,
    o.status,
    oi.id AS order_item_id,
    oi.quantity,
    oi.item_price,
    f.name AS food_name,
    f.description AS food_description,
    f.price AS food_price,
    fc.customization_type,
    fc.value AS customization_value
FROM
    "order" o
        JOIN
    orderitem oi ON o.id = oi.order_id
        JOIN
    food f ON oi.food_id = f.id
        LEFT JOIN
    ordercustomization oc ON oi.id = oc.order_item_id
        LEFT JOIN
    foodcustomization fc ON oc.food_customization_id = fc.id
WHERE
    o.id = $1
ORDER BY
    oi.id;
