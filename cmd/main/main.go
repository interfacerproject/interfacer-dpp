package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/database"
	"github.com/interfacerproject/interfacer-dpp/internal/handler"
	"github.com/interfacerproject/interfacer-dpp/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	storage.InitMinio()

	// Ensure MongoDB indexes
	dbClient, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	database.EnsureIndexes(dbClient)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // This properly allows all origins without credentials conflict
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "did-sign", "did-pk", "x-user-id"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // Must be false when allowing all origins
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/dpp", handler.CreateDPP)
	router.GET("/dpp/:id", handler.GetDPP)
	router.PUT("/dpp/:id", handler.UpdateDPP)
	router.DELETE("/dpp/:id", handler.DeleteDPP)
	router.PUT("/dpp/:id/status", handler.UpdateDPPStatus)
	router.POST("/dpp/:id/attachments", handler.AddAttachment)
	router.DELETE("/dpp/:id/attachments/:attachmentId", handler.DeleteAttachment)
	router.GET("/dpps", handler.GetAllDPPs)
	router.POST("/upload", handler.UploadFile)
	router.GET("/file/:id", handler.GetFile)

	router.Run(":8080")
}
