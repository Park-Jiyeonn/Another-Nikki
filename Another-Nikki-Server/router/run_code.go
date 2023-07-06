package router

import (
	"Another-Nikki/router/model"
	service "Another-Nikki/service/codeService"
	"Another-Nikki/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func RunCode(c *gin.Context) {
	var code model.Code
	err := c.BindJSON(&code)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	ID := uuid.NewString()
	defer service.DeleteFile(ID)

	if code.Lang == "c++" {
		if !service.CppRun(c, &code, ID) {
			return
		}
	} else if code.Lang == "python" {
		if !service.PythonRun(c, &code, ID) {
			return
		}
	} else if code.Lang == "java" {
		if !service.JavaRun(c, &code, ID) {
			return
		}
	} else if code.Lang == "go" {
		if !service.GoRun(c, &code, ID) {
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": "暂不支持这种语言呢～",
		})
		return
	}

	ret, ok := service.GetRetInFile(c, ID)
	if !ok {
		return
	}

	ans := strings.Split(ret, "\n")
	n := len(ans)
	util.SendResp(c, 200, gin.H{
		"state":         "success",
		"message":       strings.Join(ans[:n-5], "\n"),
		"cpu_time_used": ans[n-5],
		"memory_used":   ans[n-3],
		"exit_code":     ans[n-2],
		"result":        ans[n-1],
	}, "编译运行成功～")
}

func Judge(c *gin.Context) {
	var code model.Code
	err := c.BindJSON(&code)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	ID := uuid.NewString()
	defer service.DeleteFile(ID)

	if code.Lang == "c++" {
		if !service.CppJudge(c, &code, ID) {
			return
		}
	} else if code.Lang == "python" {
		if !service.PythonJudge(c, &code, ID) {
			return
		}
	} else if code.Lang == "java" {
		if !service.JavaJudge(c, &code, ID) {
			return
		}
	} else if code.Lang == "go" {
		if !service.GoJudge(c, &code, ID) {
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": "暂不支持这种语言呢～",
		})
		return
	}

	ret, ok := service.GetRetInFile(c, ID)
	if !ok {
		return
	}

	ans := strings.Split(ret, "\n")
	n := len(ans)
	util.SendResp(c, 200, gin.H{
		"state":         "success",
		"cpu_time_used": ans[n-5],
		"memory_used":   ans[n-3],
		"exit_code":     ans[n-2],
		"message":       ans[n-1],
	}, "编译运行成功～")
}
