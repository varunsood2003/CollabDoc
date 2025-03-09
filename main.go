package main

import (
	"net/http"

	"doc-editor/models"

	"github.com/gin-gonic/gin"
)

var documents = []models.Document{}

func main() {
	r := gin.Default()

	r.POST("/documents", func(c *gin.Context) {
		var doc models.Document
		if err := c.ShouldBindJSON(&doc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		documents = append(documents, doc)
		c.JSON(http.StatusCreated, doc)
	})

	r.GET("/documents", func(c *gin.Context) {
		c.JSON(http.StatusOK, documents)
	})

	r.GET("/documents/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, doc := range documents {
			if doc.ID == id {
				c.JSON(http.StatusOK, doc)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
	})

	r.DELETE("/documents/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, doc := range documents {
			if doc.ID == id {
				documents = append(documents[:i], documents[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Document deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
	})

	r.Run(":8080")
}
