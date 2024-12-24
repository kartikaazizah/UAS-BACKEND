package controllers

import (
	"context"
	"net/http"
	"BEKEN_UAS_PRAK/database"
	"BEKEN_UAS_PRAK/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllModul handles GET requests to fetch all modul
func GetAllModul(c *gin.Context) {
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var modul []models.Modul
	if err = cursor.All(context.Background(), &modul); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, modul)
}

// CreateModul handles POST requests to create a new modul
func CreateModul(c *gin.Context) {
	var modul models.Modul

	// Bind data JSON dari request ke struct Modul
	if err := c.ShouldBindJSON(&modul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat ObjectID baru untuk modul
	modul.ID = primitive.NewObjectID()

	// Ambil koleksi modul
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")

	// Simpan data ke MongoDB
	_, err := collection.InsertOne(context.Background(), modul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan data yang berhasil disimpan
	c.JSON(http.StatusOK, modul)
}



// GetModulByID handles GET requests to fetch a single modul by ID
func GetModulByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")

	var modul models.Modul
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&modul)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modul not found"})
		return
	}

	c.JSON(http.StatusOK, modul)
}

func UpdateModul(c *gin.Context) {
    id := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var modul models.Modul
    if err := c.ShouldBindJSON(&modul); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    collection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")
    _, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": modul})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update modul"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Modul updated successfully"})
}

// DeleteModul handles DELETE requests to delete a modul by ID
func DeleteModul(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Modul deleted successfully"})
}
