-- name: GetPlayer :one
SELECT * FROM player
WHERE username = $1 LIMIT 1;
