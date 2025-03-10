package handlers

import (
	"doc-editor/config"
	"doc-editor/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {
	var doc models.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&doc)
	c.JSON(http.StatusCreated, doc)
}

func GetDocuments(c *gin.Context) {
	var docs []models.Document
	config.DB.Find(&docs)
	c.JSON(http.StatusOK, docs)
}

func UpdateDocument(c *gin.Context) {
	docID := c.Param("id")
	var doc models.Document
	if err := config.DB.First(&doc, "id = ?", docID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Save(&doc)
	c.JSON(http.StatusOK, doc)
}

func DeleteDocument(c *gin.Context) {
	docID := c.Param("id")
	config.DB.Delete(&models.Document{}, "id = ?", docID)
	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
