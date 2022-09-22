package api

import (
	"database/sql"
	"net/http"
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type updatePlayerNameRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (server *Server) createPlayer(ctx *gin.Context) {

	player, err := server.store.CreatePlayer(ctx, sql.NullString{String: "", Valid: true})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func (server *Server) updatePlayerName(ctx *gin.Context) {
	var req updatePlayerNameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePlayerNameParams{
		ID:       1,
		Username: sql.NullString{String: req.Username, Valid: true},
	}

	player, err := server.store.UpdatePlayerName(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}
