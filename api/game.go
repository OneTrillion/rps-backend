package api

import (
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type addNewRpsChoiceRequest struct {
	RpsChoice int64 `uri:"rps_choice" binding:"required"`
}

func (server *Server) addNewRpsChoice(ctx *gin.Context) {
	var req addNewRpsChoiceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currentPlayerId, err := server.store.GetPlayerById(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
	}

	ctx.JSON(http.StatusOK, choice)
}
