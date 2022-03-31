package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func def(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Server working as intended")
}
func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}
func test(c *gin.Context) {
	startTime := time.Now()
	total := 10000000000
	sum := 0
	for i := 0; i < total; i++ {
		sum++
	}
	if sum == total {
		runTime := time.Since(startTime)
		c.JSON(200, runTime.Milliseconds())
	} else {
		c.JSON(500, "Server Error")
	}
}

func main() {
	router := gin.Default()
	router.GET("/", def)
	router.GET("/ping", ping)
	router.GET("/test", test)

	router.Run("localhost:4568")
}
