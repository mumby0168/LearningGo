package handlers

import (
	"../newsfeed"
	"github.com/gin-gonic/gin"
)

func GetNewsfeed(repo *newsfeed.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		result := repo.GetAll()
		context.JSON(200, result)
	}
}
