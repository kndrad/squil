-- name: AllDogs :many
SELECT *
FROM dogs
ORDER BY name;
-- name: CreateDog :one
INSERT INTO dogs (name, breed)
VALUES ($1, $2)
RETURNING *;
-- name: ReadDog :one
SELECT *
FROM dogs
WHERE name = $1
LIMIT 1;
-- name: UpdateDog :one
UPDATE dogs
SET name = COALESCE($2, name),
    breed = COALESCE($3, breed)
WHERE id = $1
RETURNING *;
-- name: DeleteDog :exec
DELETE FROM dogs
WHERE name = $1;
