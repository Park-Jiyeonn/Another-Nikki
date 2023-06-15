package main

import (
	"Another-Nikki/dal"
	"Another-Nikki/mw"
	"Another-Nikki/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dal.Init()

	r.Use(mw.Cors())

	blogGroup := r.Group("/api/blog")
	{
		blogGroup.GET("/get_all_blogs", router.GetAllBlogs)
		blogGroup.GET("/get_last_seven_blogs", router.GetLastSevenBlog)
		blogGroup.POST("/create_blog", router.PostBlog)
		blogGroup.GET("/get_random_blog", router.GetRandomBlog)
	}
	
	r.Run(":8888")
}
