package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.StaticFile("/application.js", "./public/application.js")
	router.StaticFile("/application.css", "./public/application.css")

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	return router
}

func main() {
	setupRouter().Run(":8080")
}
