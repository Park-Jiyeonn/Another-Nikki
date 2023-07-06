package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

var problem dal.Problem

func PostProblem(c *gin.Context) {
	var newProblem model.Problem
	err := c.BindJSON(&newProblem)
	if util.HandleError(c, err, "参数填写错误") {
		return
	}

	err = problem.Create(newProblem.Name, newProblem.Content)
	if util.HandleError(c, err, "数据库创建记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "success")
}

func GetProblemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("ID"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := problem.GetInfoById(id)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetProblemByPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := problem.GetList(page)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func UpdateProblem(c *gin.Context) {
	var newProblem model.ProblemUpdate
	err := c.BindJSON(&newProblem)
	if util.HandleError(c, err, "参数填写错误") {
		return
	}
	err = problem.UpdateProblem(newProblem.ID, newProblem.Name, newProblem.Content)
	if util.HandleError(c, err, "数据库更新记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "更新成功")
}
