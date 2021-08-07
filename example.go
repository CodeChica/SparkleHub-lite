package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.StaticFile("/index.html", "./public/index.html")

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})
	router.Run(":8080")
}
