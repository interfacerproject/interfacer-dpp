package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/handler"
)

func main() {
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

	router.Run(":8080")
}
