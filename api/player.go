package api

import (
	"fmt"
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

	playerHealth, err := server.store.GetPlayerHealth(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if playerHealth == 25 {
		// TODO lose function
	}

	arg := db.UpdatePlayerHealthParams{
		ID:     int64(currentPlayerId),
		Health: playerHealth - 25,
	}

	health, err := server.store.UpdatePlayerHealth(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, health)
}

// This is for the ult power, will reset players health to 100hp
func (server *Server) resetPlayerHealth(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdatePlayerHealthParams{
		ID:     int64(currentPlayerId),
		Health: 100,
	}

	newHealth, err := server.store.UpdatePlayerHealth(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newHealth)

}

// Takes the players ult value
func (server *Server) getPlayerUltMeter(ctx *gin.Context) {
	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	playersUltValue, err := server.store.GetPlayersUlt(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playersUltValue)

}

// Will increase the players ult bar every time the player wins
func (server *Server) increasePlayerUltMeter(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	playerUlt, err := server.store.GetPlayersUlt(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if playerUlt >= 100 {
		// Function that will tell ult is already 100%
		err := fmt.Errorf("ult is full")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePlayersUltParams{
		ID:       int64(currentPlayerId),
		UltMeter: playerUlt + 25,
	}

	newUltMeter, err := server.store.UpdatePlayersUlt(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUltMeter)

}

func (server *Server) resetUltMeter(ctx *gin.Context) {
	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdatePlayersUltParams{
		ID:       int64(currentPlayerId),
		UltMeter: 0,
	}

	reseter, err := server.store.UpdatePlayersUlt(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// SET HEALTH TO 100% resetPlayerHealth function

	server.resetPlayerHealth(ctx)

	ctx.JSON(http.StatusOK, reseter)
}

func (server *Server) getScore(ctx *gin.Context) {
	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	score, err := server.store.GetPlayerScore(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, score)
}

func (server *Server) updateScore(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updateScore, err := server.store.UpdateScore(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updateScore)
}

func (server *Server) getUsername(ctx *gin.Context) {

	// takes the current player id and checks for error
	currentPlayerId, err := server.currentPlayerId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	username, err := server.store.GetUsername(ctx, int64(currentPlayerId))
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, username)

}
