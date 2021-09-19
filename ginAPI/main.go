package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	router := gin.Default()
	router.GET("hello", helloWorld)
	router.Run(":80")
}
