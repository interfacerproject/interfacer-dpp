package main

import (
	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/handler"
)

func main() {
	router := gin.Default()

	router.POST("/dpp", handler.CreateDPP)
	router.GET("/dpp/:id", handler.GetDPP)
	router.PUT("/dpp/:id", handler.UpdateDPP)
	router.DELETE("/dpp/:id", handler.DeleteDPP)
	router.GET("/dpps", handler.GetAllDPPs)

	router.Run("localhost:8080")
}
