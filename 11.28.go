package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

improt(
	"web"
	"net/http"
)
func main() {
	router:=gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong！！~")
	})
	router.GET("/echo", func(c *gin.Context) {
		message := c.Query("message")
		c.String(http.StatusOK, message)
	})
	router.Run(`:8080`)
}
