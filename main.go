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

func (self *SparkleRepository) Save(sparkle Sparkle) {
	self.Sparkles = append(self.Sparkles, sparkle)
}

func createSparkleFrom(body string) Sparkle {
	items := strings.SplitAfterN(body, " ", 2)
	return Sparkle{
		Sparklee: User{Name: items[0]},
		Reason:   items[1],
	}
}

var db SparkleRepository = SparkleRepository{
	Sparkles: []Sparkle{},
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": db.Sparkles})
	})
	router.POST("/sparkles", func(context *gin.Context) {
		db.Save(createSparkleFrom(context.PostForm("body")))
		context.Redirect(http.StatusFound, "/")
	})
	return router
}

func main() {
	setupRouter().Run(":8080")
}
