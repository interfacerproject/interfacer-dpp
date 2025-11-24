package handler

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/auth"
	"github.com/interfacerproject/interfacer-dpp/internal/database"
	"github.com/interfacerproject/interfacer-dpp/internal/model"
	"github.com/oklog/ulid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/interfacerproject/interfacer-dpp/internal/storage"
	"github.com/minio/minio-go/v7"
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

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Request Body: %s\n", string(body))

	zenroomData := auth.ZenroomData{
		Gql:            b64.StdEncoding.EncodeToString(body),
		EdDSASignature: c.Request.Header.Get("did-sign"),
		EdDSAPublicKey: c.Request.Header.Get("did-pk"),
	}

	if err := zenroomData.VerifyDid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "DID verification failed", "details": err.Error()})
		return
	}

	if err := zenroomData.IsAuth(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed", "details": err.Error()})
		return
	}

	var dpp model.DigitalProductPassport
	if err := json.Unmarshal(body, &dpp); err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving documents", "details": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &dpps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding documents"})
		return
	}

	c.JSON(http.StatusOK, dpps)
}

func GetFile(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	ctx := context.Background()

	// List objects with the fileID prefix to find the file with any extension
	objectCh := storage.MinioClient.ListObjects(ctx, storage.BucketName, minio.ListObjectsOptions{
		Prefix: fileID,
	})

	var objectName string
	for object := range objectCh {
		if object.Err != nil {
			log.Printf("Error listing objects: %v", object.Err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding file"})
			return
		}
		if object.Key != "" {
			objectName = object.Key
			break
		}
	}

	if objectName == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Get the object from MinIO
	obj, err := storage.MinioClient.GetObject(ctx, storage.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Error getting object: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving file"})
		return
	}
	defer obj.Close()

	// Get object info for content type
	objInfo, err := obj.Stat()
	if err != nil {
		log.Printf("Error getting object info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving file info"})
		return
	}

	// Set content type and stream the file
	c.Header("Content-Type", objInfo.ContentType)
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", objectName))

	if _, err := io.Copy(c.Writer, obj); err != nil {
		log.Printf("Error streaming file: %v", err)
		return
	}
}

func UploadFile(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required", "details": err.Error()})
		return
	}
	defer file.Close()

	hasher := sha256.New()
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(io.MultiWriter(hasher, buffer), file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing file"})
		return
	}
	checksum := hex.EncodeToString(hasher.Sum(nil))

	signature := c.Request.Header.Get("did-sign")
	publicKey := c.Request.Header.Get("did-pk")

	if signature == "" || publicKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication headers (did-sign, did-pk)"})
		return
	}

	zenroomData := auth.ZenroomData{
		Gql:            checksum,
		EdDSASignature: signature,
		EdDSAPublicKey: publicKey,
	}

	if err := zenroomData.VerifyDid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "DID verification failed", "details": err.Error()})
		return
	}

	if err := zenroomData.IsAuth(); err != nil {
		log.Printf("Auth failed for file upload: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed", "details": err.Error()})
		return
	}

	ext := filepath.Ext(header.Filename)
	fileID := ulid.Make().String()
	objectName := fmt.Sprintf("%s%s", fileID, ext)
	contentType := header.Header.Get("Content-Type")

	ctx := context.Background()
	_, err = storage.MinioClient.PutObject(ctx, storage.BucketName, objectName, buffer, int64(buffer.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Printf("MinIO upload error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to storage"})
		return
	}

	fileURL := fmt.Sprintf("%s/file/%s", storage.ServerURL, fileID)

	attachment := model.Attachment{
		ID:          fileID,
		FileName:    header.Filename,
		ContentType: contentType,
		URL:         fileURL,
		Size:        header.Size,
		Checksum:    checksum,
		UploadedAt:  time.Now(),
	}

	c.JSON(http.StatusCreated, attachment)
}
