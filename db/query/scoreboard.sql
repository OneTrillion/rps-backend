-- name: AddNewScoreToScoreboard :one
INSERT INTO scoreboard (player_id, username, score)
VALUES ($1, $2, $3) 
RETURNING *;
