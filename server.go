package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func NewServer(sparkles *[]Sparkle) *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.Use(cors.Default())

	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": sparkles})
	})

	router.POST("/sparkles.json", func(context *gin.Context) {
		var form map[string]string
		if err := context.BindJSON(&form); err != nil {
			context.String(http.StatusUnprocessableEntity, err.Error())
		} else {
			sparkle, err := NewSparkle(form["body"])
			if err == nil {
				*sparkles = append(*sparkles, *sparkle)
				context.JSON(http.StatusCreated, gin.H{
					"reason":   sparkle.Reason,
					"sparklee": sparkle.Sparklee,
				})
			} else {
				context.String(http.StatusUnprocessableEntity, err.Error())
			}
		}
	})

	return router
}
