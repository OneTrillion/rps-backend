-- name: GetPlayer :one
SELECT * FROM player
WHERE username = $1 LIMIT 1;

-- name: CreatePlayer :one
INSERT INTO player (
    username
) VALUES (
    $1
) RETURNING *;