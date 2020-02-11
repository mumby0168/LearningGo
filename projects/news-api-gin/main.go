package main

import (
	"./handlers"
	"./newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	newsfeed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.GET("/newsfeed", handlers.GetNewsfeed(newsfeed))
	r.POST("/newsfeed", handlers.PostNewsFeed(newsfeed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
