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

	// SCOREBOARD FUNCTIONS
	router.POST("/scoreboard", server.addNewScore)
	router.GET("/scoreboard", server.listHigestScores)

	// PLAYER FUNCTIONS
	router.POST("/player", server.createPlayer)
	router.POST("/player/:id", server.updatePlayerName)
	router.GET("/score", server.getScore)

	// HEALTH FUNCTIONS
	router.GET("/get/health", server.getPlayerHealth)
	router.POST("/decrease/health", server.decreasePlayerHealth)
	router.POST("/use/ult", server.resetPlayerHealth)

	// ULT FUNCTIONS
	router.GET("/get/ult", server.getPlayerUltMeter)
	router.POST("/increase/ult", server.increasePlayerUltMeter)
	router.POST("/reset/ult", server.resetUltMeter)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
