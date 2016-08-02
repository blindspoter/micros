package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bluesprint(router *gin.Engine) {
	router.GET("/ping", handlerPing)
}

func handlerPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
