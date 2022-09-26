-- name: AddNewRpsChoice :one
INSERT INTO choice (player_id, rps_choice)
VALUES ($1, $2) RETURNING *;

-- name: GetPlayerChoice :one
SELECT rps_choice
FROM choice
WHERE player_id = $1
LIMIT 1;