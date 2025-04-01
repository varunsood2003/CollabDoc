package routes

import (
	"doc-editor/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		AuthRoutes(api)
		DocRoutes(api)
	}

	router.GET("/ws", handlers.HandleWebSocket)
}

func AuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}
}

func DocRoutes(router *gin.RouterGroup) {
	docs := router.Group("/documents")
	{
		docs.POST("/", handlers.CreateDocument)
		docs.GET("/", handlers.GetDocuments)
		docs.PUT("/:id", handlers.UpdateDocument)
		docs.DELETE("/:id", handlers.DeleteDocument)
	}
}
