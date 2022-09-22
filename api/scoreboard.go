package api

import (
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type addNewScoreRequest struct {
	PlayerID int64  `json:"player_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Score    int32  `json:"score" binding:"required"`
}

func (server *Server) addNewScore(ctx *gin.Context) {
	var req addNewScoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddNewScoreToScoreboardParams{
		PlayerID: req.PlayerID,
		Username: req.Username,
		Score:    req.Score,
	}

	newScore, err := server.store.AddNewScoreToScoreboard(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newScore)
}
