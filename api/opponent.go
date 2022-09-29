package api

import (
	"net/http"
	"rps-backend/util"

	"github.com/gin-gonic/gin"
)

// Creates opponent with a random generated name from util
func (server *Server) createOpponent(ctx *gin.Context) {

	opponentName := util.RandomName()

	opponent, err := server.store.CreateOpponent(ctx, opponentName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, opponent)
}

func (server *Server) getOpponentsHealth(ctx *gin.Context) {

	health, err := server.store.GetOpponentHealth(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if health == 0 {

		server.createOpponent(ctx)
		server.updateScore(ctx)
	}

}

func (server *Server) decreaseOpponentsHealth(ctx *gin.Context) {

	_, err := server.store.DecreaseOpponentHealth(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

}
