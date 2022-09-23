package api

import (
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Will take the current players id and return it
func (server *Server) currentPlayerId(ctx *gin.Context) (int, error) {
	returningPlayerId, err := server.store.GetPlayerById(ctx)

	var playerId int = int(returningPlayerId.(int64))
	return playerId, err
}

// Creates an empty player profile with the name player
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
	// takes the current player id and checks for error
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

// Takes the current players health
func (server *Server) getPlayerHealth(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	playerHealth, err := server.store.GetPlayerHealth(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playerHealth)

}

func (server *Server) decreasePlayerHealth(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdatePlayerHealthParams{
		ID:     int64(currentPlayerId),
		Health: -25,
	}

	health, err := server.store.UpdatePlayerHealth(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, health)
}
