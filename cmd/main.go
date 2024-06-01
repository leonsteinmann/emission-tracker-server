package main

import (
	"emission-tracker-server/internal/config"
	"emission-tracker-server/internal/database"
	"emission-tracker-server/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB(config.GetDBConfig())
	defer db.Close()

	router := gin.Default()

	// Define routes
	router.GET("/records", handler.GetRecords(db))
	router.GET("/records/:id", handler.GetRecordByID(db))
	router.POST("/records", handler.PostRecord(db))

	router.Run("localhost:8080")
}
