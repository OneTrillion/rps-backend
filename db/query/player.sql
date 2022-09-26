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
WHERE id = $1
RETURNING *;

-- name: UpdatePlayerHealth :one
UPDATE player
SET health = $2
WHERE id = $1
RETURNING *;

-- name: GetPlayerHealth :one
SELECT health FROM player
WHERE id = $1 LIMIT 1;

-- name: GetPlayerById :one
SELECT MAX(id) FROM player LIMIT 1;

-- name: UpdatePlayersUlt :one
UPDATE player
SET ult_meter = $2
WHERE id = $1
RETURNING *;

-- name: GetPlayersUlt :one
SELECT ult_meter FROM player
WHERE id = $1 LIMIT 1;


-- name: GetPlayerScore :one
SELECT score FROM player
WHERE id = $1 LIMIT 1;

-- name: UpdateScore :one
UPDATE player
SET score = score + 1
WHERE id = $1
RETURNING *;
