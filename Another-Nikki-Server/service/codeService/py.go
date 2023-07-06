package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PythonRun(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "py.py")
	if util.HandleError(c, err, "写入 py.py 或 data.in 失败") {
		return false
	}
	cmd := fmt.Sprintf("docker run --rm -m 256m --name py_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i data.in -o data.out -i data.in python3 -u py.py >> data.out 2>&1'", ID, ID)
	err = RunCommand(cmd)
	ret, _ := GetRetInFile(c, ID)
	if strings.Contains(ret, "Error") {
		ans := strings.Split(ret, "\n")
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": "run fail: \n" + strings.Join(ans[:len(ans)-3], "\n"),
		})
		return false
	}

	return !util.HandleError(c, err, "run faild")
}

func PythonJudge(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "py.py")
	if util.HandleError(c, err, "写入 py.py 或 data.in 失败") {
		return false
	}
	cmd := fmt.Sprintf("docker run --rm -m 256m --name py_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i ../onlineJudge/problemData/%s/input.in -a ../onlineJudge/problemData/%s/output.out -o out.out -t 2000 python3 -u py.py >> data.out 2>&1'", ID, ID, code.ProblemName, code.ProblemName)
	err = RunCommand(cmd)
	ret, _ := GetRetInFile(c, ID)
	if strings.Contains(ret, "Error") {
		ans := strings.Split(ret, "\n")
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": "run fail: \n" + strings.Join(ans[:len(ans)-3], "\n"),
		})
		return false
	}

	return !util.HandleError(c, err, "run faild")
}
