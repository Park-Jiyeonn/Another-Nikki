package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Python(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "py.py")
	if util.HandleError(c, err, "写入 py.py 或 data.in 失败") {
		return false
	}
	cmd := fmt.Sprintf("docker run --rm -m 256m --name py_run-%s -v $(pwd)/code/tmp-%s:/dox oj:1 sh -c './../calc/calc2 python3 -u py.py < data.in > data.out 2>&1'", ID,ID)
	err = RunCommand(cmd)
	ret, _ := GetRetInFile(c, ID)
	if strings.Contains(ret, "Error") {
		ans := strings.Split(ret, "\n")
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": "run fail: \n" + strings.Join(ans[:len(ans) - 3], "\n"),
		})
		return false
	}

	return !util.HandleError(c, err, "run faild")
}