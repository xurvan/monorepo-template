-- name: ListUsers :many
SELECT *
FROM users
ORDER BY create_time
LIMIT sqlc.arg('pageSize')
OFFSET sqlc.arg('pageToken');

-- name: GetUser :one
SELECT *
FROM users
WHERE id = sqlc.arg('id');

-- name: CreateUser :one
INSERT INTO users (name, age)
VALUES (sqlc.arg('name'), sqlc.arg('age'))
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;
