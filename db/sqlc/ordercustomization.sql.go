// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ordercustomization.sql

package db

import (
	"context"
)

const addCustomizationToOrderItem = `-- name: AddCustomizationToOrderItem :one
INSERT INTO ordercustomization (order_item_id, food_customization_id)
VALUES ($1, $2)
    RETURNING order_item_id, food_customization_id
`

type AddCustomizationToOrderItemParams struct {
	OrderItemID         int32 `json:"order_item_id"`
	FoodCustomizationID int32 `json:"food_customization_id"`
}

func (q *Queries) AddCustomizationToOrderItem(ctx context.Context, arg AddCustomizationToOrderItemParams) (Ordercustomization, error) {
	row := q.db.QueryRowContext(ctx, addCustomizationToOrderItem, arg.OrderItemID, arg.FoodCustomizationID)
	var i Ordercustomization
	err := row.Scan(&i.OrderItemID, &i.FoodCustomizationID)
	return i, err
}

const deleteOrderItemCustomization = `-- name: DeleteOrderItemCustomization :exec
DELETE FROM ordercustomization
WHERE order_item_id = $1 AND food_customization_id = $2
`

type DeleteOrderItemCustomizationParams struct {
	OrderItemID         int32 `json:"order_item_id"`
	FoodCustomizationID int32 `json:"food_customization_id"`
}

func (q *Queries) DeleteOrderItemCustomization(ctx context.Context, arg DeleteOrderItemCustomizationParams) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItemCustomization, arg.OrderItemID, arg.FoodCustomizationID)
	return err
}

const getOrderItemCustomizations = `-- name: GetOrderItemCustomizations :many
SELECT fc.customization_type, fc.value
FROM ordercustomization oc
         JOIN foodcustomization fc ON oc.food_customization_id = fc.id
WHERE oc.order_item_id = $1
`

type GetOrderItemCustomizationsRow struct {
	CustomizationType string `json:"customization_type"`
	Value             string `json:"value"`
}

func (q *Queries) GetOrderItemCustomizations(ctx context.Context, orderItemID int32) ([]GetOrderItemCustomizationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderItemCustomizations, orderItemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrderItemCustomizationsRow
	for rows.Next() {
		var i GetOrderItemCustomizationsRow
		if err := rows.Scan(&i.CustomizationType, &i.Value); err != nil {
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
