package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jdotc2/whisper/go/config"
	"github.com/jdotc2/whisper/go/routes"
)

func main() {
	// Config.Test()
	config.Connect()

	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}