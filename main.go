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

type InMemoryDatabase struct {
	Sparkles []Sparkle
}

func (db *InMemoryDatabase) Save(sparkle Sparkle) {
	db.Sparkles = append(db.Sparkles, sparkle)
}

func squish(item string) string {
	return strings.Trim(item, " ")
}

func NewUser(name string) User {
	return User{Name: name}
}

func NewSparkle(body string) Sparkle {
	items := strings.SplitAfterN(body, " ", 2)
	username := squish(items[0])
	reason := squish(items[1])

	return Sparkle{
		Sparklee: NewUser(username),
		Reason:   reason,
	}
}

var db InMemoryDatabase = InMemoryDatabase{
	Sparkles: []Sparkle{},
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/sparkles.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "sparkles/index.tmpl", gin.H{
			"sparkles": db.Sparkles,
		})
	})
	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": db.Sparkles})
	})
	router.POST("/sparkles", func(context *gin.Context) {
		db.Save(NewSparkle(context.PostForm("body")))
		context.Redirect(http.StatusFound, "/")
	})
	return router
}

func main() {
	setupRouter().Run(":8080")
}
