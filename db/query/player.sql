-- name: GetPlayer :one
SELECT * FROM player
WHERE username = $1 LIMIT 1;

-- name: CreatePlayer :one
INSERT INTO player (
    username
) VALUES (
    $1
) RETURNING *;

-- name: UpdatePlayerName :one
UPDATE player
SET username = $2
WHERE username = $1
RETURNING *;

-- name: UpdatePlayerHealth :one
UPDATE player
SET health = $2
WHERE username = $1
RETURNING *;
