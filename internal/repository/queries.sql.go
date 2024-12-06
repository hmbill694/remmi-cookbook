// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package repository

import (
	"context"
)

const getAllRecipes = `-- name: GetAllRecipes :many
SELECT
    id, user_id, title, description, created_at
from
    recipes
`

func (q *Queries) GetAllRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.Query(ctx, getAllRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchForRecipes = `-- name: SearchForRecipes :many
SELECT
    id, user_id, title, description, created_at
from
    recipes
where
    title ILIKE '%' || $1::text || '%'
    or description ILIKE '%' || $1::text || '%'
`

func (q *Queries) SearchForRecipes(ctx context.Context, search string) ([]Recipe, error) {
	rows, err := q.db.Query(ctx, searchForRecipes, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}