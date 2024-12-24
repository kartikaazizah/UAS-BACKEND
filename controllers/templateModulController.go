package controllers

import (
	"BEKEN_UAS_PRAK/database"
	"BEKEN_UAS_PRAK/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTemplateModul creates a new template_modul
func CreateTemplateModul(c *gin.Context) {
	var template models.TemplateModul

	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template.ID = primitive.NewObjectID()

	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")
	_, err := collection.InsertOne(context.Background(), template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, template)
}

// GetAllTemplateModul retrieves all template_modul entries
func GetAllTemplateModul(c *gin.Context) {
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var templates []models.TemplateModul
	if err := cursor.All(context.Background(), &templates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, templates)
}

// GetTemplateModulByID retrieves a single template_modul by ID
func GetTemplateModulByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")
	var template models.TemplateModul
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&template)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template modul not found"})
		return
	}

	c.JSON(http.StatusOK, template)
}

// UpdateTemplateModul updates an existing template_modul
func UpdateTemplateModul(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var template models.TemplateModul
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": template})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update template modul"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template modul updated successfully"})
}

// DeleteTemplateModul deletes a template_modul by ID
func DeleteTemplateModul(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete template modul"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template modul deleted successfully"})
}
