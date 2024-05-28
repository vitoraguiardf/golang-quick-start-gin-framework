package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	routerAdmin := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"vitor": "rotiv",
	}))
	routerAdmin.GET("/index", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Admin is here!")
	})
	return router
}
