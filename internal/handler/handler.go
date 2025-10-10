package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/database"
	"github.com/interfacerproject/interfacer-dpp/internal/model"
	"github.com/oklog/ulid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getCollection() (*mongo.Collection, error) {
	client, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	return database.GetCollection(client), nil
}

func CreateDPP(c *gin.Context) {
	dppCollection, err := getCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dpp model.DigitalProductPassport
	if err := c.BindJSON(&dpp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dpp.ID = ulid.Make()

	_, err = dppCollection.InsertOne(ctx, dpp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"insertedID": dpp.ID.String()})
}

func GetDPP(c *gin.Context) {
	dppCollection, err := getCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objId, err := ulid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dpp model.DigitalProductPassport
	err = dppCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&dpp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving document"})
		return
	}

	c.JSON(http.StatusOK, dpp)
}

func UpdateDPP(c *gin.Context) {
	dppCollection, err := getCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objId, err := ulid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dpp model.DigitalProductPassport
	if err := c.BindJSON(&dpp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{"$set": dpp}

	result, err := dppCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating document"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"matchedCount": result.MatchedCount, "modifiedCount": result.ModifiedCount})
}

func DeleteDPP(c *gin.Context) {
	dppCollection, err := getCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objId, err := ulid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := dppCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting document"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedCount": result.DeletedCount})
}

func GetAllDPPs(c *gin.Context) {
	dppCollection, err := getCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var dpps []model.DigitalProductPassport
	cursor, err := dppCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving documents"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &dpps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding documents"})
		return
	}

	c.JSON(http.StatusOK, dpps)
}
