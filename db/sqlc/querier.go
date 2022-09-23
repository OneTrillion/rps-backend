// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AddNewScoreToScoreboard(ctx context.Context, arg AddNewScoreToScoreboardParams) (Scoreboard, error)
	CreatePlayer(ctx context.Context, username string) (Player, error)
	GetPlayer(ctx context.Context, username string) (Player, error)
	GetPlayerById(ctx context.Context) (interface{}, error)
	ListHigestScores(ctx context.Context) ([]ListHigestScoresRow, error)
	UpdatePlayerHealth(ctx context.Context, arg UpdatePlayerHealthParams) (Player, error)
	UpdatePlayerName(ctx context.Context, arg UpdatePlayerNameParams) (Player, error)
}

var _ Querier = (*Queries)(nil)
