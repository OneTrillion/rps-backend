-- name: CreateOpponent :one
INSERT INTO opponent (
    opponent_name
) VALUES (
    $1
) RETURNING *;

-- name: GetOpponentHealth :one
SELECT health
FROM opponent 
WHERE id=(
    SELECT max(id) FROM opponent
    )
LIMIT 1;

-- name: DecreaseOpponentHealth :one
UPDATE opponent 
SET health = health - 25
WHERE id=(
    SELECT max(id) FROM opponent
    )
RETURNING health;

