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
		blogGroup.GET("/get_all_blogs", router.GetAllComments)
		blogGroup.GET("/get_last_seven_blogs", router.GetLastSevenComment)
		blogGroup.POST("/create_blog", router.PostComment)
		blogGroup.GET("/get_random_blog", router.GetRandomComment)
	}

	runCode := r.Group("/api/runcode")
	{
		runCode.POST("/run", router.RunCode)
		runCode.POST("/judge", router.Judge)
		runCode.GET("/default_code", router.DefaultCode)
	}

	user := r.Group("/api/user")
	{
		user.POST("/login", router.Login)
		user.POST("/register", router.Register)
	}

	article := r.Group("/api/article")
	{
		article.GET("", router.GetArticleById)
		article.GET("/articles", router.GetArticleByPage)
		// 这里要鉴权
		article.POST("/update", router.UpdateArticle).Use(mw.JwtAuth())
		article.POST("", router.PostArticle).Use(mw.JwtAuth())
	}

	problem := r.Group("/api/problem")
	{
		problem.GET("/problems", router.GetProblemByPage)
		problem.GET("", router.GetProblemById)
		// 这里要鉴权
		problem.POST("", router.PostProblem).Use(mw.JwtAuth())
		problem.POST("/update", router.UpdateProblem).Use(mw.JwtAuth())
	}

	r.Run(":8888")
}
