package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	// Controllers
	"github.com/jdotc2/whisper/go/controllers"
)

// Routes Function
func Routes(router *gin.Engine) {
	router.Use(cors.Default())
	router.GET("/", welcome)
	router.NoRoute(notFound)
	router.GET("/users", controllers.GetAllUsers)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Whisper API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}