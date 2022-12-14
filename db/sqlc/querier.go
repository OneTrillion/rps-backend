// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AddNewRpsChoice(ctx context.Context, arg AddNewRpsChoiceParams) (Game, error)
	AddNewScoreToScoreboard(ctx context.Context, arg AddNewScoreToScoreboardParams) (Scoreboard, error)
	CreateOpponent(ctx context.Context, opponentName string) (Opponent, error)
	CreatePlayer(ctx context.Context, username string) (Player, error)
	DecreaseOpponentHealth(ctx context.Context) (int32, error)
	FinalizeGame(ctx context.Context, arg FinalizeGameParams) (Player, error)
	GetOpponentHealth(ctx context.Context) (int32, error)
	GetPlayer(ctx context.Context, username string) (Player, error)
	GetPlayerById(ctx context.Context) (interface{}, error)
	GetPlayerChoice(ctx context.Context) (int32, error)
	GetPlayerHealth(ctx context.Context, id int64) (int32, error)
	GetPlayerScore(ctx context.Context, id int64) (int32, error)
	GetPlayersUlt(ctx context.Context, id int64) (int32, error)
	GetUsername(ctx context.Context, id int64) (string, error)
	ListHigestScores(ctx context.Context) ([]ListHigestScoresRow, error)
	UpdatePlayerHealth(ctx context.Context, arg UpdatePlayerHealthParams) (Player, error)
	UpdatePlayersUlt(ctx context.Context, arg UpdatePlayersUltParams) (Player, error)
	UpdateScore(ctx context.Context, id int64) (Player, error)
}

var _ Querier = (*Queries)(nil)
