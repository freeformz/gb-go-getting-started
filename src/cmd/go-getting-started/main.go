package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	version = ""
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, version)

	})

	router.Run(":" + port)
}
