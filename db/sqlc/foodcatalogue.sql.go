// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: foodcatalogue.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO foodcatalogue (
    name
) VALUES (
             $1
         ) RETURNING id, name
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Foodcatalogue, error) {
	row := q.db.QueryRowContext(ctx, createCategory, name)
	var i Foodcatalogue
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM foodcatalogue
WHERE name = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, name)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name FROM foodcatalogue
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Foodcatalogue, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Foodcatalogue
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE foodcatalogue
SET name = $2
WHERE id = $1
    RETURNING id, name
`

type UpdateCategoryParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Foodcatalogue, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.Name)
	var i Foodcatalogue
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
