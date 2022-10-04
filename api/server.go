package api

import (
	db "rps-backend/db/sqlc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	//INSERT api path here example POST GET etc...

	// GAME FUNCTIONS
	router.POST("/game/:rps_choice", server.addNewRpsChoice)
	router.GET("/status", server.ifLost)

	// SCOREBOARD FUNCTIONS
	router.POST("/scoreboard", server.addNewScore)
	router.GET("/scoreboard", server.listHigestScores)

	// PLAYER FUNCTIONS
	router.POST("/player", server.createPlayer)
	router.GET("/score", server.getScore)
	router.PUT("/add/score", server.updateScore)
	router.GET("/get/username", server.getUsername)

	// OPPONENT FUNCTIONS
	router.POST("/opponent", server.createOpponent)
	router.GET("/opponent/hp", server.getOpponentsHealth)
	router.PUT("/decrease/opponent", server.decreaseOpponentsHealth)

	// POST PLAYER TO SCOREBOARD AND ADD NAME
	router.POST("/finalize", server.finalizeGame)

	// HEALTH FUNCTIONS
	router.GET("/get/health", server.getPlayerHealth)
	router.PUT("/decrease/health", server.decreasePlayerHealth)
	router.PUT("/use/ult", server.resetPlayerHealth)

	// ULT FUNCTIONS
	router.GET("/get/ult", server.getPlayerUltMeter)
	router.PUT("/increase/ult", server.increasePlayerUltMeter)
	router.PUT("/reset/ult", server.resetUltMeter)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
