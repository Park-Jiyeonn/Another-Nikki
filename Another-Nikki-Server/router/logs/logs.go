package logs

import (
	"Another-Nikki/dal"
	"Another-Nikki/mw"
	"Another-Nikki/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var log dal.Log

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetLogCount(c *gin.Context) {
	sum, err := log.GetLogCount()
	if util.HandleError(c, err, "出错了捏，查不到日志") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"onlineJudge": 200,
		"data": gin.H{
			"sum": sum,
		},
		"message": "success",
	})
}

func GetPageQue(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := log.GetPageQue(page)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	for i := range ret {
		ret[i].Response = ret[i].Response[:min(50, len(ret[i].Response))]
	}
	c.JSON(http.StatusOK, gin.H{
		"onlineJudge": 200,
		"data": gin.H{
			"logs": ret,
		},
		"message": "success",
	})
}

func GetVisitTime(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	visitTime := mw.GetVisitTime(int64(userID))
	c.JSON(http.StatusOK, gin.H{
		"onlineJudge": 200,
		"data": gin.H{
			"sum_visit_time":   visitTime >> 32,
			"today_visit_time": visitTime & (1<<32 - 1),
		},
		"message": "success",
	})
}
