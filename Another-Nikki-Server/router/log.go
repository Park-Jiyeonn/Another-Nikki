package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/util"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetLogCount(c *gin.Context) {
	sum, err := dal.GetLogCount(context.Background())
	if util.HandleError(c, err, "出错了捏，查不到日志") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sum": sum,
	})
}

func GetPageQue(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := dal.GetPageQue(context.Background(), page)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	for i := range ret {
		ret[i].Response = ret[i].Response[:min(50, len(ret[i].Response))]
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}