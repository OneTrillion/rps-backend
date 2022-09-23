package api

import (
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) currentPlayerId(ctx *gin.Context) (int, error) {
	returningPlayerId, err := server.store.GetPlayerById(ctx)

	var playerId int = int(returningPlayerId.(int64))
	return playerId, err
}

func (server *Server) createPlayer(ctx *gin.Context) {

	player, err := server.store.CreatePlayer(ctx, "player")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

type updatePlayerNameRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (server *Server) updatePlayerName(ctx *gin.Context) {
	var req updatePlayerNameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdatePlayerNameParams{
		ID:       int64(currentPlayerId),
		Username: req.Username,
	}

	player, err := server.store.UpdatePlayerName(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func (server *Server) decreasePlayerHealth(ctx *gin.Context) {

	arg := db.UpdatePlayerHealthParams{
		ID:     1,
		Health: -25,
	}

	health, err := server.store.UpdatePlayerHealth(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, health)
}
