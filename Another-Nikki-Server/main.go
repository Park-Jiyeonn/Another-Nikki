package main

import (
	"Another-Nikki/mw"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(mw.Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": "快要有前端啦",
		})
	})
	r.POST("/github-webhook", GithubWebhook)
	r.Run(":8888")
}
