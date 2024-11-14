package api

import "github.com/gin-gonic/gin"

// create a type to hold the server

type Server struct {
	store  *db.Store
	router *gin.Engine
}
