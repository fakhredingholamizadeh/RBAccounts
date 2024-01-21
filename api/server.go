package api

import (
	db "github.com/flfreecode/rbaccounts/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our user services.
type Server struct {
	store *db.Store

	router *gin.Engine // Gin Engine Pointer
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	return server

}

// Start HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
