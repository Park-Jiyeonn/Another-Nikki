package router

import (
	"Another-Nikki/dal"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Content string `json:"content"`
}

func PostBlog(c *gin.Context) {
	var blog Blog
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据库创建记录失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func GetAllBlogs(c *gin.Context) {
	ret, err := dal.GetBlogList(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据库查询记录失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}

func GetLastSevenBlog(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请填写合适的参数" + err.Error(),
		})
		return
	}
	ret, err := dal.GetLastSevenBlog(context.Background(), int64(num))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据库查询记录失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}

func GetRandomBlog(c *gin.Context) {
	ret, err := dal.GetRandomBlog(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据库查询记录失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}