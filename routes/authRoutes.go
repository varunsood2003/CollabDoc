package routes

import (
	"doc-editor/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
}
