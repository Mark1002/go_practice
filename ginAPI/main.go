package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": os.Getenv("MESSAGE"),
	})
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("hello", helloWorld)
	return router
}

func main() {
	router := setUpRouter()
	router.Run(":80")
}
