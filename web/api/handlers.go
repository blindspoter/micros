package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"micros/models"
)

func Bluesprint(router *gin.Engine) {
	router.GET("/ping", handlerPing)
	// for test
	router.GET("/user/:id/name", handlerUserName)
}

func handlerPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func handlerUserName(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)
	name, err := models.GetUserNameById(int64(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "get user name failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": name})
}
