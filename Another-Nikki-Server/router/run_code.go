package router

import (
	"Another-Nikki/router/model"
	service "Another-Nikki/service/codeService"
	"Another-Nikki/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RunCode(c *gin.Context) {
	var code model.Code
	err := c.BindJSON(&code)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	ID := time.Now().UnixNano()
	defer service.DeleteFile(ID)
	
	if code.Lang == "c++" {
		if !service.Cpp(c, &code, ID) {
			return
		}
	} else if code.Lang == "python" {
		if !service.Python(c, &code, ID) {
			return
		}
	} else if code.Lang == "java" {
		if !service.Java(c, &code, ID) {
			return
		}
	} else if code.Lang == "go" {
		if !service.Go(c, &code, ID) {
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state":"error",
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
	c.JSON(http.StatusOK, gin.H{
		"state":"success",
		"message":strings.Join(ans[:n-3], "\n"),
		"cpu_time_used":ans[n-3],
		"memory_used":ans[n-2],
		"exit_code":ans[n-1],
	})
}
