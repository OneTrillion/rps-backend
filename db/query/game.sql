-- name: AddNewRpsChoice :one
INSERT INTO game (player_id, rps_choice)
VALUES ($1, $2) RETURNING *;

-- name: GetPlayerChoice :one
SELECT rps_choice 
FROM game 
WHERE MAX (id)
LIMIT 1;