package api

import (
	"net/http"
	db "rps-backend/db/sqlc"
	"rps-backend/util"

	"github.com/gin-gonic/gin"
)

// ONCLICK "START GAME"
// Start new game, create new player, start timer
func (server *Server) startGame(ctx *gin.Context) {

	// Create new opponent
	server.createOpponent(ctx)

	// Create new player
	server.createPlayer(ctx)

	// Start timer

}

type addNewRpsChoiceRequest struct {
	RpsChoice int64 `uri:"rps_choice" binding:"required"`
}

// ONCLICK pick menu "confirm" add to database
func (server *Server) addNewRpsChoice(ctx *gin.Context) {
	var req addNewRpsChoiceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currentPlayerId, err := server.store.GetPlayerById(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var i interface{} = currentPlayerId
	s := i.(int64)

	arg := db.AddNewRpsChoiceParams{
		PlayerID:  s,
		RpsChoice: int32(req.RpsChoice),
	}

	choice, err := server.store.AddNewRpsChoice(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, choice)
}

func (server *Server) compareChoices(ctx *gin.Context) {
	choice, err := server.store.GetPlayerChoice(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	computerChoice := util.RandomRpsChoice()

	if choice == computerChoice {
		// New round
		ctx.JSON(http.StatusOK, "Same choice")
	}

	// Win
	if (choice == 1 && computerChoice == 3) ||
		(choice == 2 && computerChoice == 1) ||
		(choice == 3 && computerChoice == 2) {

		// Add point to ult
		server.increasePlayerUltMeter(ctx)
		// Deal damage to computer
		server.decreaseOpponentsHealth(ctx)
		server.getOpponentsHealth(ctx)
		// New round

	} else {

		// Take damage
		server.decreasePlayerHealth(ctx)

		// New round

	}

}

// Will check if player has lost, if lost = true
func (server *Server) ifLost(ctx *gin.Context) {

	// Update frontend with Health

	currentPlayerId, err := server.store.GetPlayerById(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var i interface{} = currentPlayerId
	id := i.(int64)

	health, err := server.store.GetPlayerHealth(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if health == 0 {
		ctx.JSON(http.StatusOK, true)
		return
	}
	ctx.JSON(http.StatusOK, false)

}
