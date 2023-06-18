package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/util"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Content string `json:"content"`
}

func PostBlog(c *gin.Context) {
	var blog Blog
	err := c.ShouldBindJSON(&blog)
	if util.HandleError(c, err, "参数错误") {
		return
	}
	if blog.Content == "" {
		c.JSON(http.StatusOK, gin.H{
			"message":"不可以上传空的留言哦~",
		})
		return
	}
	if strings.TrimSpace(blog.Content) == "" {
		c.JSON(http.StatusOK, gin.H{
			"message":"全都是空格不行的呢~",
		})
		return
	}
	err = dal.CreateBolg(context.Background(), blog.Content)
	if util.HandleError(c, err, "数据库创建记录失败") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func GetAllBlogs(c *gin.Context) {
	ret, err := dal.GetBlogList(context.Background())
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}

func GetLastSevenBlog(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := dal.GetLastSevenBlog(context.Background(), int64(num))
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}

func GetRandomBlog(c *gin.Context) {
	ret, err := dal.GetRandomBlog(context.Background())
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
	time.Sleep(time.Millisecond * 100)
}