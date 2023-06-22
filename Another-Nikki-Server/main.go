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

	logGroup := r.Group("/api/log")
	{
		logGroup.GET("/get_page_que", router.GetPageQue)
		logGroup.GET("/count", router.GetLogCount)
	}

	r.Use(mw.Logger())

	blogGroup := r.Group("/api/blog")
	{
		blogGroup.GET("/get_all_blogs", router.GetAllBlogs)
		blogGroup.GET("/get_last_seven_blogs", router.GetLastSevenBlog)
		blogGroup.POST("/create_blog", router.PostBlog)
		blogGroup.GET("/get_random_blog", router.GetRandomBlog)
	}

	runCode := r.Group("/api/runcode")
	{
		runCode.POST("", router.RunCode)
		runCode.GET("/default_code", router.DefaultCode)
	}
	
	r.Run(":8888")
}
