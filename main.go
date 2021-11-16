package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monkrus/gingormfun.git/model"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "test"})
	})
	model.ConnectDatabase() // new

	r.Run()
}
