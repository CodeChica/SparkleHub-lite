package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Squish(item string) string {
	return strings.Trim(item, " ")
}

type Sparkle struct {
	Sparklee string `json:"sparklee"`
	Reason   string `json:"reason"`
}

func NewSparkle(body string) Sparkle {
	items := strings.SplitAfterN(body, " ", 2)
	username := Squish(items[0])
	reason := Squish(items[1])

	return Sparkle{
		Sparklee: username,
		Reason:   reason,
	}
}

func setupRouter(sparkles *[]Sparkle) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/sparkles.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "sparkles/index.tmpl", gin.H{
			"sparkles": sparkles,
		})
	})

	router.GET("/sparkles.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"sparkles": sparkles})
	})

	router.POST("/sparkles", func(context *gin.Context) {
		sparkle := NewSparkle(context.PostForm("body"))
		*sparkles = append(*sparkles, sparkle)

		context.Redirect(http.StatusFound, "/")
	})
	return router
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func listenAddress() string {
	return ":" + port()
}

func main() {
	sparkles := []Sparkle{}
	log.Fatal(setupRouter(&sparkles).Run(listenAddress()))
}
