-- CreateOpponent :one
INSERT INTO opponent (
    opponent_name
) VALUES (
    $1
) RETURNING *;

-- GetOpponentHealth :one
SELECT health 
FROM opponent
WHERE MAX (id) LIMIT 1;

-- DecreaseOpponentHealth :one
UPDATE opponent 
SET health = health - 25
WHERE MAX (id) 
RETURNING *;

