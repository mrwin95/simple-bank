package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mrwin95/simple-bank/db/sqlc"
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

	server.router = router

	return server
}
