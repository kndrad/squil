// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: dog_query.sql

package shelter

import (
	"context"
)

const allDogs = `-- name: AllDogs :many
SELECT id, name, breed, created_at, updated_at
FROM dogs
ORDER BY name
`

func (q *Queries) AllDogs(ctx context.Context) ([]Dog, error) {
	rows, err := q.db.Query(ctx, allDogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dog
	for rows.Next() {
		var i Dog
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Breed,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const createDog = `-- name: CreateDog :one
INSERT INTO dogs (name, breed)
VALUES ($1, $2)
RETURNING id, name, breed, created_at, updated_at
`

type CreateDogParams struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func (q *Queries) CreateDog(ctx context.Context, arg CreateDogParams) (Dog, error) {
	row := q.db.QueryRow(ctx, createDog, arg.Name, arg.Breed)
	var i Dog
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteDog = `-- name: DeleteDog :exec
DELETE FROM dogs
WHERE name = $1
`

func (q *Queries) DeleteDog(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, deleteDog, name)
	return err
}

const readDog = `-- name: ReadDog :one
SELECT id, name, breed, created_at, updated_at
FROM dogs
WHERE name = $1
LIMIT 1
`

func (q *Queries) ReadDog(ctx context.Context, name string) (Dog, error) {
	row := q.db.QueryRow(ctx, readDog, name)
	var i Dog
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateDog = `-- name: UpdateDog :one
UPDATE dogs
SET name = COALESCE($2, name),
    breed = COALESCE($3, breed)
WHERE id = $1
RETURNING id, name, breed, created_at, updated_at
`

type UpdateDogParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func (q *Queries) UpdateDog(ctx context.Context, arg UpdateDogParams) (Dog, error) {
	row := q.db.QueryRow(ctx, updateDog, arg.ID, arg.Name, arg.Breed)
	var i Dog
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
