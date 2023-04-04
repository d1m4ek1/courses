package main

import (
	"courses/internal/database"
	"courses/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var db database.Context

	ctxdb, err := db.Init()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	gin.SetMode(gin.ReleaseMode)

	handle := handlers.Server{Router: gin.Default(), ProductionMode: false}
	handle.InitServer(ctxdb)
}
