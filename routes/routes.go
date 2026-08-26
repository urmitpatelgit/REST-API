package routes

import (
	"net/http"

	"example.com/rest-api/middlewares"
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

	// Option 1 : Add Authenticate as 1st function to execute
	//
	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	//Option 2 : Group
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Executed before any requests in group is fulfilled
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
