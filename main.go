package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	hub := newHub()
	go hub.run()

	router := gin.New()
	router.LoadHTMLFiles("client.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "client.html", nil)
	})

	router.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	router.Run(":3000")
}
