-- name: CreateProject :one
INSERT INTO projects (id, name, description, image_url, github_url, tech_stack, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetProjectByID :one
SELECT * FROM projects WHERE id = $1;

-- name: GetProjects :many
SELECT * FROM projects ORDER BY created_at DESC;

-- name: UpdateProject :one
UPDATE projects
SET name = $2,
    description = $3,
    image_url = $4,
    github_url = $5,
    tech_stack = $6,
    updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;