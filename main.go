package main

import (
	"doc-editor/config"
	"doc-editor/handlers"
	"doc-editor/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)
	go handlers.BroadcastMessages()

	r.Run(":8080")
}
