package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

type Sparkle struct {
	Sparklee User
	Reason   string
}

type SparkleRepository struct {
	Sparkles []Sparkle
}

func (self *SparkleRepository) Insert(sparklee string, reason string) {
	self.Sparkles = append(self.Sparkles, Sparkle{
		Sparklee: User{Name: sparklee},
		Reason:   reason,
	})
}

var db SparkleRepository = SparkleRepository{
	Sparkles: []Sparkle{},
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": db.Sparkles})
	})
	router.POST("/sparkles", func(context *gin.Context) {
		body := context.PostForm("body")
		items := strings.SplitAfterN(body, " ", 2)
		db.Insert(items[0], items[1])
		context.Redirect(http.StatusFound, "/")
	})
	return router
}

func main() {
	setupRouter().Run(":8080")
}
