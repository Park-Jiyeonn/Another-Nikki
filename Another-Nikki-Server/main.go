package main

import (
	"Another-Nikki/dal"
	"Another-Nikki/mw"
	"Another-Nikki/router/Article"
	"Another-Nikki/router/Comment"
	"Another-Nikki/router/Problem"
	"Another-Nikki/router/RunCode"
	"Another-Nikki/router/User"
	"Another-Nikki/router/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dal.Init()

	r.Use(mw.Cors())

	// log 不走 log 记录哈哈哈
	logGroup := r.Group("/api/log")
	// log 要鉴权
	logGroup.Use(mw.JwtAuth(true))
	{
		logGroup.GET("/get_page_que", logs.GetPageQue)
		logGroup.GET("/count", logs.GetLogCount)
	}

	r.Use(mw.Logger())

	go mw.SetTodayTime(1)
	go mw.SetTodayTime(2)
	logGroup.GET("/visit_time", logs.GetVisitTime)

	// blog 要鉴权
	commentGroup := r.Group("/api/blog")
	commentGroup.Use(mw.JwtAuth(false))
	{
		commentGroup.GET("/get_all_blogs", Comment.GetAllComments)
		commentGroup.GET("/get_last_seven_blogs", Comment.GetLastSevenComment)
		commentGroup.POST("/create_blog", Comment.PostComment)
		commentGroup.GET("/get_random_blog", Comment.GetRandomComment)
	}

	runCodeGroup := r.Group("/api/runcode")
	{
		runCodeGroup.POST("/run", RunCode.RunCode)
		runCodeGroup.POST("/judge", RunCode.Judge)
		runCodeGroup.GET("/default_code", RunCode.DefaultCode)
	}

	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", User.Login)
		userGroup.POST("/register", User.Register)
	}

	articleGroup := r.Group("/api/article")
	{
		articleGroup.GET("", Article.GetArticleById)
		articleGroup.GET("/articles", Article.GetArticleByPage)
		// 这里要鉴权
		articleGroup.POST("/update", Article.UpdateArticle).Use(mw.JwtAuth(true))
		articleGroup.POST("", Article.PostArticle).Use(mw.JwtAuth(true))
	}

	problemGroup := r.Group("/api/problem")
	{
		problemGroup.GET("/problems", Problem.GetProblemByPage)
		problemGroup.GET("", Problem.GetProblemById)
		// 这里要鉴权
		problemGroup.POST("", Problem.PostProblem).Use(mw.JwtAuth(true))
		problemGroup.POST("/update", Problem.UpdateProblem).Use(mw.JwtAuth(true))
	}

	r.POST("/api/upload", Comment.UploadImage)

	r.Run(":8888")
}
