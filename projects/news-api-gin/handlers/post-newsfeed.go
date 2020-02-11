package handlers

import (
	"../newsfeed"
	"github.com/gin-gonic/gin"
)

type addNewsFeedRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func PostNewsFeed(repo newsfeed.IAdder) gin.HandlerFunc {
	return func(context *gin.Context) {

		requestBody := addNewsFeedRequest{}
		context.Bind(&requestBody)

		repo.Add(newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		})

		context.Status(200)
	}
}
