package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var comment dal.Comment

func PostBlog(c *gin.Context) {
	var blog model.Blog
	err := c.ShouldBindJSON(&blog)
	if util.HandleError(c, err, "参数错误") {
		return
	}
	if blog.Content == "" {
		util.SendResp(c, 404, nil, "不可以上传空的留言哦~")
		return
	}
	if strings.TrimSpace(blog.Content) == "" {
		util.SendResp(c, 404, nil, "全都是空格不行的呢～")
		return
	}
	err = comment.CreateBolg(blog.Content)
	if util.HandleError(c, err, "数据库创建记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "success")
}
func GetAllBlogs(c *gin.Context) {
	ret, err := comment.GetBlogList()
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetLastSevenBlog(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := comment.GetLastSevenBlog(int64(num))
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetRandomBlog(c *gin.Context) {
	ret, err := comment.GetRandomBlog()
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
	time.Sleep(time.Millisecond * 100)
}