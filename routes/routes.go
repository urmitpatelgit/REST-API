package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Define a simple GET endpoint
	server.GET("/ping", func(context *gin.Context) {
		// Return JSON response
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1   events/105
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("events/:id", deleteEvent)

}
