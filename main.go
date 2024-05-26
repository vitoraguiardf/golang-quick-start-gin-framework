package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}
