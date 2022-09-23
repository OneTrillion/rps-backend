-- name: AddNewScoreToScoreboard :one
INSERT INTO scoreboard (player_id, username, score)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: ListHigestScores :many
SELECT username, score 
FROM scoreboard
ORDER BY score DESC
LIMIT 10;