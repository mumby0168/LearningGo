package handlers

import (
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, map[string]string{
		"Alive": "True",
	})
}
