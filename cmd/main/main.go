package main

import (
	"courses/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	handle := handlers.Server{Router: gin.Default(), ProductionMode: false}
	handle.InitServer()
}
