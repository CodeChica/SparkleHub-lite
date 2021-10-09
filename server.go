package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func NewServer(sparkles *[]Sparkle) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.Use(cors.Default())

	router.GET("/sparkles.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "sparkles/index.tmpl", gin.H{
			"sparkles": sparkles,
		})
	})

	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": sparkles})
	})

	router.POST("/sparkles", func(context *gin.Context) {
		sparkle, err := NewSparkle(context.PostForm("body"))
		if err == nil {
			*sparkles = append(*sparkles, *sparkle)

			context.Redirect(http.StatusFound, "/")
		} else {
			context.String(http.StatusUnprocessableEntity, "")
		}
	})
	return router
}
