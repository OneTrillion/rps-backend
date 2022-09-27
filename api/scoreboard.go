package api

import (
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) addNewScore(ctx *gin.Context) {

	// Takes the players id
	currentPlayerId, err := server.store.GetPlayerById(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// Convert currentPlayerId from interface to int
	var i interface{} = currentPlayerId
	id := i.(int64)

	// Take username with the id
	username, err := server.store.GetUsername(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// Takes players score with the id
	score, err := server.store.GetPlayerScore(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.AddNewScoreToScoreboardParams{
		PlayerID: id,
		Username: username,
		Score:    score,
	}

	newScore, err := server.store.AddNewScoreToScoreboard(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newScore)
}

func (server *Server) listHigestScores(ctx *gin.Context) {
	scores, err := server.store.ListHigestScores(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scores)
}
