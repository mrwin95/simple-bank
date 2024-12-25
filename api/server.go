package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mrwin95/simple_bank/db/sqlc"
)

// create a type to hold the server

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new server

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router

	return server
}

// Start server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
