package main

import (
	"github.com/ayushpal11/wav-flac-converter/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static HTML page
	router.StaticFile("/", "./static/index.html")

	// Initialize WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		server.HandleWebSocket(c.Writer, c.Request)
	})

	// Start the server
	router.Run(":8080")
}
