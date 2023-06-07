package router

import (
	"Another-Nikki/dal"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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
	}
	err = dal.CreateBolg(context.Background(), blog.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据库创建记录失败：" + err.Error(),
		})
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
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}
