package main

import (
	"log"

	"doc-editor/config"
	"doc-editor/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	routes.RegisterAuthRoutes(r)
	// routes.RegisterDocumentRoutes(r)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
