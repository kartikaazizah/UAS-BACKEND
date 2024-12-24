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

// GetAllUsers handles GET requests to fetch all users
func GetAllUsers(c *gin.Context) {
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("user")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User

	// Bind data JSON dari body request ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate ObjectID baru untuk user
	user.ID = primitive.NewObjectID()

	// Ambil koleksi "user"
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("user")

	// Simpan data user ke MongoDB
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan respons dengan data user yang berhasil disimpan
	c.JSON(http.StatusOK, user)
}

// GetUserByID handles GET requests to fetch a single user by ID
func GetUserByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("user")

	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserModules(c *gin.Context) {
	// Ambil ID user dari parameter URL
	userID := c.Param("id")

	// Convert ID user ke ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ambil koleksi user_modul
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("user_modul")

	// Query untuk mencari modul berdasarkan user_id
	filter := bson.M{"user_id": objID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	// Decode hasil query ke dalam slice UserModul
	var userModules []models.UserModul
	if err := cursor.All(context.Background(), &userModules); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ambil koleksi modul
	modulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("modul")

	// Cari detail modul berdasarkan modul_id
	var modulDetails []models.Modul
	for _, userModul := range userModules {
		var modul models.Modul
		err := modulCollection.FindOne(context.Background(), bson.M{"_id": userModul.Modul_ID}).Decode(&modul)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch modul details"})
			return
		}
		modulDetails = append(modulDetails, modul)
	}

	// Kembalikan daftar modul yang dimiliki user
	c.JSON(http.StatusOK, userModules)
}

func UpdateJenisUser(c *gin.Context) {
	// Ambil ID user dari parameter URL
	userID := c.Param("id")

	// Parse ID user ke ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ambil jenis_user baru dari body request
	var body struct {
		JenisUser primitive.ObjectID `json:"jenis_user"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update jenis_user di koleksi user
	userCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("user")
	_, err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"jenis_user": body.JenisUser}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Hapus semua modul lama milik user di koleksi user_modul
	userModulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("user_modul")
	_, err = userModulCollection.DeleteMany(context.Background(), bson.M{"user_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user modules"})
		return
	}

	// Ambil template modul untuk jenis_user baru dari koleksi template_modul
	templateModulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("template_modul")
	cursor, err := templateModulCollection.Find(context.Background(), bson.M{"id_jenis_user": body.JenisUser})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch template modules"})
		return
	}
	defer cursor.Close(context.Background())

	// Tambahkan modul baru ke user_modul berdasarkan template
	var templateModules []models.UserModul
	if err := cursor.All(context.Background(), &templateModules); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template modules"})
		return
	}

	for _, tmpl := range templateModules {
		newUserModul := models.UserModul{
			ID:       primitive.NewObjectID(),
			User_ID:  objID,
			Modul_ID: tmpl.Modul_ID,
		}
		_, err := userModulCollection.InsertOne(context.Background(), newUserModul)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new user module"})
			return
		}
	}

	// Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Jenis user updated successfully and modules assigned"})
}

func AddModulToUser(c *gin.Context) {
	userID := c.Param("id")
	objUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var body struct {
		ModulID primitive.ObjectID `json:"modul_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("user_modul")

	newUserModul := models.UserModul{
		ID:      primitive.NewObjectID(),
		User_ID:  objUserID,
		Modul_ID: body.ModulID,
	}

	_, err = userModulCollection.InsertOne(context.Background(), newUserModul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add modul to user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Modul added to user successfully"})
}

func UpdateUserModul(c *gin.Context) {
	userID := c.Param("id")
	modulID := c.Param("modul_id")

	objUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	objModulID, err := primitive.ObjectIDFromHex(modulID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid modul ID"})
		return
	}

	var body struct {
		NewModulID primitive.ObjectID `json:"new_modul_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("user_modul")

	_, err = userModulCollection.UpdateOne(
		context.Background(),
		bson.M{"user_id": objUserID, "modul_id": objModulID},
		bson.M{"$set": bson.M{"modul_id": body.NewModulID}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user modul"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User modul updated successfully"})
}

func DeleteUserModul(c *gin.Context) {
	userID := c.Param("id")
	modulID := c.Param("modul_id")

	objUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	objModulID, err := primitive.ObjectIDFromHex(modulID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid modul ID"})
		return
	}

	userModulCollection := database.Client.Database("SSO_UAS_BEKEN").Collection("user_modul")

	_, err = userModulCollection.DeleteOne(context.Background(), bson.M{"user_id": objUserID, "modul_id": objModulID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user modul"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User modul deleted successfully"})
}



// DeleteUser handles DELETE requests to delete a user by ID
func DeleteUser(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.Client.Database("SSO_UAS_BEKEN").Collection("user")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
