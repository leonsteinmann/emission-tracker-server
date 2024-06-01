package main

import (
	"emission-tracker-server/internal/config"
	"emission-tracker-server/internal/database"
	"emission-tracker-server/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB(config.GetDBConfig())
	defer db.Close()

	router := gin.Default()

	// Serve static files (HTML, CSS, JS)
	router.LoadHTMLGlob("web/templates/*")

	// Define route to serve HTML page for the dashboard
	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// GET
	router.GET("/records", handler.GetRecords(db))
	router.GET("/records/:id", handler.GetRecordByID(db))
	// POST
	router.POST("/records", handler.PostRecord(db))
	// DELETE
	router.DELETE("/records/:id", handler.DeleteRecordByID(db))

	router.Run("localhost:8080")

}
