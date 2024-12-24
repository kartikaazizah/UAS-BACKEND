package controllers

import (
	"context"
	"log"
	"net/http"
	"BEKEN_UAS_PRAK/database"
	"BEKEN_UAS_PRAK/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllKategori retrieves all categories
func GetAllKategori(c *gin.Context) {
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("kategori")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var kategori []models.Kategori
	if err = cursor.All(context.Background(), &kategori); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// CreateKategori creates a new category
func CreateKategori(c *gin.Context) {
    var kategori models.Kategori
    if err := c.ShouldBindJSON(&kategori); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println("Data diterima dari Postman:", kategori)

    collection := database.Client.Database("SSO_UAS_BEKEN").Collection("kategori")
    kategori.ID = primitive.NewObjectID()

    _, err := collection.InsertOne(context.Background(), kategori)
    if err != nil {
        log.Println("Error inserting data:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    log.Println("Data berhasil disimpan ke MongoDB:", kategori)
    c.JSON(http.StatusOK, kategori)
}

func CreateManyKategori(c *gin.Context) {
	var kategoris []models.Kategori
	if err := c.ShouldBindJSON(&kategoris); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("kategori")

	var docs []interface{}
	for _, kategori := range kategoris {
		kategori.ID = primitive.NewObjectID()
		docs = append(docs, kategori)
	}

	_, err := collection.InsertMany(context.Background(), docs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategoris inserted successfully"})
}
