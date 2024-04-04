-- name: GetUser :one
SELECT * FROM users WHERE username = $1;

-- name: AddUser :one
INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;