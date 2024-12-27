// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO "order" (
                     order_date,
                     total_price,
                     status)
VALUES (
        $1,
        $2,
        $3)
    RETURNING id, order_date, total_price, status
`

type CreateOrderParams struct {
	OrderDate  string `json:"order_date"`
	TotalPrice int32  `json:"total_price"`
	Status     string `json:"status"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.OrderDate, arg.TotalPrice, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderDate,
		&i.TotalPrice,
		&i.Status,
	)
	return i, err
}

const getFullOrderDetails = `-- name: GetFullOrderDetails :many

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
    oi.id
`

type GetFullOrderDetailsRow struct {
	OrderID            int32          `json:"order_id"`
	OrderDate          string         `json:"order_date"`
	TotalPrice         int32          `json:"total_price"`
	Status             string         `json:"status"`
	OrderItemID        int32          `json:"order_item_id"`
	Quantity           int32          `json:"quantity"`
	ItemPrice          int32          `json:"item_price"`
	FoodName           string         `json:"food_name"`
	FoodDescription    string         `json:"food_description"`
	FoodPrice          int32          `json:"food_price"`
	CustomizationType  sql.NullString `json:"customization_type"`
	CustomizationValue sql.NullString `json:"customization_value"`
}

// Special action
func (q *Queries) GetFullOrderDetails(ctx context.Context, id int32) ([]GetFullOrderDetailsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFullOrderDetails, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFullOrderDetailsRow
	for rows.Next() {
		var i GetFullOrderDetailsRow
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderDate,
			&i.TotalPrice,
			&i.Status,
			&i.OrderItemID,
			&i.Quantity,
			&i.ItemPrice,
			&i.FoodName,
			&i.FoodDescription,
			&i.FoodPrice,
			&i.CustomizationType,
			&i.CustomizationValue,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderByStatus = `-- name: GetOrderByStatus :one
SELECT id, order_date, total_price, status FROM "order"
WHERE status = $1
ORDER BY id
`

func (q *Queries) GetOrderByStatus(ctx context.Context, status string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByStatus, status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderDate,
		&i.TotalPrice,
		&i.Status,
	)
	return i, err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :one
UPDATE "order"
SET status = $2
WHERE id = $1
    RETURNING id, order_date, total_price, status
`

type UpdateOrderStatusParams struct {
	ID     int32  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrderStatus, arg.ID, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderDate,
		&i.TotalPrice,
		&i.Status,
	)
	return i, err
}