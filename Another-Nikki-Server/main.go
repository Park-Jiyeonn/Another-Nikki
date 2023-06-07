package main

import (
	"Another-Nikki/dal"
	"Another-Nikki/mw"
	"Another-Nikki/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	dal.Init()

	r.Use(mw.Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": "快要有前端啦",
		})
	})
	r.POST("/github-webhook", router.GithubWebhook)

	r.POST("/api/create_blog", router.PostBlog)
	r.GET("/api/get_all_blogs", router.GetAllBlogs)

	r.Run(":8888")
}

// test webhook 1
