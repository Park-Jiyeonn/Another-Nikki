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

func PostComment(c *gin.Context) {
	var newComment model.Comment
	err := c.ShouldBindJSON(&newComment)
	if util.HandleError(c, err, "参数错误") {
		return
	}
	if newComment.Content == "" {
		util.SendResp(c, 404, nil, "不可以上传空的留言哦~")
		return
	}
	if strings.TrimSpace(newComment.Content) == "" {
		util.SendResp(c, 404, nil, "全都是空格不行的呢～")
		return
	}
	err = comment.CreateComment(newComment.Content,
		newComment.AuthorName,
		newComment.ParentName,
		newComment.AuthorID,
		newComment.RootID,
		newComment.ParentID)
	if util.HandleError(c, err, "数据库创建记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "success")
}

func GetAllComments(c *gin.Context) {
	ret, err := comment.GetCommentList()
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetLastSevenComment(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := comment.GetLastSevenComment(int64(num))
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetRandomComment(c *gin.Context) {
	ret, err := comment.GetRandomComment()
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
	time.Sleep(time.Millisecond * 100)
}
