// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: scoreboard.sql

package db

import (
	"context"
)

const addNewScoreToScoreboard = `-- name: AddNewScoreToScoreboard :one
INSERT INTO scoreboard (player_id, username, score)
VALUES ($1, $2, $3) 
RETURNING id, player_id, username, score
`

type AddNewScoreToScoreboardParams struct {
	PlayerID int64  `json:"player_id"`
	Username string `json:"username"`
	Score    int32  `json:"score"`
}

func (q *Queries) AddNewScoreToScoreboard(ctx context.Context, arg AddNewScoreToScoreboardParams) (Scoreboard, error) {
	row := q.db.QueryRowContext(ctx, addNewScoreToScoreboard, arg.PlayerID, arg.Username, arg.Score)
	var i Scoreboard
	err := row.Scan(
		&i.ID,
		&i.PlayerID,
		&i.Username,
		&i.Score,
	)
	return i, err
}

const listHigestScores = `-- name: ListHigestScores :many
SELECT username, score 
FROM scoreboard
ORDER BY score DESC
LIMIT 10
`

type ListHigestScoresRow struct {
	Username string `json:"username"`
	Score    int32  `json:"score"`
}

func (q *Queries) ListHigestScores(ctx context.Context) ([]ListHigestScoresRow, error) {
	rows, err := q.db.QueryContext(ctx, listHigestScores)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHigestScoresRow{}
	for rows.Next() {
		var i ListHigestScoresRow
		if err := rows.Scan(&i.Username, &i.Score); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
