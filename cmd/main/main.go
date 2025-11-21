package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "did-sign", "did-pk"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/dpp", handler.CreateDPP)
	router.GET("/dpp/:id", handler.GetDPP)
	router.PUT("/dpp/:id", handler.UpdateDPP)
	router.DELETE("/dpp/:id", handler.DeleteDPP)
	router.GET("/dpps", handler.GetAllDPPs)
	router.POST("/upload", handler.UploadFile)

	router.Run(":8080")
}
