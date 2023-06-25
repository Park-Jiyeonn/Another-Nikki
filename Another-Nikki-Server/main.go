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

	// log 不走 log 记录哈哈哈
	logGroup := r.Group("/api/log")
	// log 要鉴权
	logGroup.Use(mw.JwtAuth())
	{
		logGroup.GET("/get_page_que", router.GetPageQue)
		logGroup.GET("/count", router.GetLogCount)
	}

	r.Use(mw.Logger())

	// blog 要鉴权
	blogGroup := r.Group("/api/blog")
	blogGroup.Use(mw.JwtAuth())
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

	user := r.Group("/api/user")
	{
		user.POST("/login", router.Login)
		user.POST("/register", router.Register)
	}

	article := r.Group("/api/article")
	{
		article.POST("", router.PostArticle)
		article.GET("", router.GetArticleById)
		article.GET("/articles", router.GetArticleByPage)
	}
	
	r.Run(":8888")
}
