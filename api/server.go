package api

import (
	db "rps-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//INSERT api path here example POST GET etc...
	router.GET("/scoreboard", server.listHigestScores)
	router.POST("/scoreboard", server.addNewScore)
	router.POST("/player", server.createPlayer)
	router.POST("/player/:id", server.createPlayer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
